package pgstore

type HealthcheckRepository struct {
	store *Store
}

func (r *HealthcheckRepository) Check() error {
	_, err := r.store.db.Exec(`SELECT 1`)

	return err
}
