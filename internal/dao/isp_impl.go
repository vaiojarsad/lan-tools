package dao

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewISPDaoImpl() ISPDao {
	return &databaseSqlIspDaoImpl{}
}

type databaseSqlIspDaoImpl struct {
}

func (d *databaseSqlIspDaoImpl) Insert(e *entities.ISP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO isp(code, name, public_ip_getter_type, public_ip_getter_cfg, " +
		"public_ip, public_ip_modified) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	pubGetCfg, err := json.Marshal(e.PublicIpGetterCfg)
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(e.Code, e.Name, e.PublicIpGetterType, pubGetCfg, e.PublicIp,
		e.PublicIpModTime.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlIspDaoImpl) GetByCode(code string) (*entities.ISP, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT id, name, public_ip_getter_type, public_ip_getter_cfg, public_ip, " +
		"public_ip_modified FROM isp WHERE code = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var id int64
	var name, publicIpGetterType, publicIpGetterCfgStr, publicIp, publicIpModTimeStr string

	err = stmt.QueryRow(code).Scan(&id, &name, &publicIpGetterType, &publicIpGetterCfgStr, &publicIp,
		&publicIpModTimeStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	publicIpModTime, err := time.Parse(time.RFC3339, publicIpModTimeStr)
	if err != nil {
		return nil, err
	}

	var publicIpGetterCfg map[string]string
	err = json.Unmarshal([]byte(publicIpGetterCfgStr), &publicIpGetterCfg)
	if err != nil {
		return nil, err
	}

	return entities.NewISP(id, code, name, publicIpGetterType, publicIpGetterCfg, publicIp, publicIpModTime), nil
}

func (d *databaseSqlIspDaoImpl) UpdatePublicIpInfo(e *entities.ISP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("UPDATE isp SET public_ip = ?, public_ip_modified = ? WHERE id = ?")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.PublicIp, e.PublicIpModTime.Format(time.RFC3339), e.StorageId())
	if err != nil {
		return err
	}

	return nil
}
