package isppublicip

import (
	"database/sql"
	"time"

	"github.com/vaiojarsad/cloudflare-tools/internal/database"
	entities "github.com/vaiojarsad/cloudflare-tools/internal/entities/isppublicip"
	"github.com/vaiojarsad/cloudflare-tools/internal/utils"
)

func New() Dao {
	return &databaseSqlDaoImpl{}
}

type databaseSqlDaoImpl struct {
}

func (d *databaseSqlDaoImpl) GetByISP(isp string) (*entities.ISPPublicIP, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT id, ip, modified FROM isp_public_ip WHERE isp = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var id int64
	var ip, modified string

	err = stmt.QueryRow(isp).Scan(&id, &ip, &modified)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, modified)
	if err != nil {
		t = time.Time{}
	}

	return entities.New(id, isp, ip, t), nil
}

func (d *databaseSqlDaoImpl) Insert(e *entities.ISPPublicIP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO isp_public_ip(isp, ip, modified) VALUES (?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.ISP, e.IP, e.Modified.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDaoImpl) UpdateIpInfo(e *entities.ISPPublicIP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("UPDATE isp_public_ip SET ip = ?, modified = ? WHERE id = ?")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.IP, e.Modified.Format(time.RFC3339), e.StorageId())
	if err != nil {
		return err
	}

	return nil
}
