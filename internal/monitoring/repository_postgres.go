package monitoring

import "context"

func (r *repository) Count() (int, error) {
	var total int

	query := `SELECT COUNT(*)FROM monitoring_edc`

	err := r.db.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) FindAll(
	limit int,
	offset int,
) ([]Monitoring, error) {

	rows, err := r.db.Query(`
		SELECT
		    id,
			tgl,
			kdcab,
			cabang,
			kdtk,
			nama,
			station,
			cek,
			ip,
			edc_bca,
			edc_mandiri,
			edc_mti,
			edc_mdr_mti
		FROM monitoring_edc
		ORDER BY tgl DESC
		LIMIT $1 OFFSET $2
		`,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Monitoring

	for rows.Next() {
		var item Monitoring

		err := rows.Scan(
			&item.ID,
			&item.Tgl,
			&item.Kdcab,
			&item.Cabang,
			&item.Kdtk,
			&item.Nama,
			&item.Station,
			&item.Cek,
			&item.IP,
			&item.EDCBCA,
			&item.EDCMandiri,
			&item.EDCMTI,
			&item.EDCMDRMTI,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

func (r *repository) BulkInsert(
	ctx context.Context,
	data []Monitoring,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	query := `
	INSERT INTO monitoring_edc
	(
		tgl,
		kdcab,
		cabang,
		kdtk,
		nama,
		station,
		cek,
		ip,
		edc_bca,
		edc_mandiri,
		edc_mti,
		edc_mdr_mti
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
	)
	`

	for _, item := range data {

		_, err := tx.ExecContext(
			ctx,
			query,
			item.Tgl,
			item.Kdcab,
			item.Cabang,
			item.Kdtk,
			item.Nama,
			item.Station,
			item.Cek,
			item.IP,
			item.EDCBCA,
			item.EDCMandiri,
			item.EDCMTI,
			item.EDCMDRMTI,
		)

		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}
