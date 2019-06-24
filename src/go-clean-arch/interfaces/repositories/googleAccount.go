package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go-clean-arch/domain/models"
	"go-clean-arch/domain/repositories"
)

type GoogleAccountRepository struct {
	Db *sqlx.DB
}

func NewGoogleAccountRepository(Db *sqlx.DB) repositories.GoogleAccountRepository {
	return &GoogleAccountRepository{Db}
}

func (googleAccountRepo *GoogleAccountRepository) FindByClientCustomerIds(ids []int) ([]models.GoogleAccount, error) {
	var googleAccounts []models.GoogleAccount

	query := `
		SELECT
			gac.google_account_refresh_token,
			gac.google_account_user_nme,
			gaw.google_adwords_id,
			gaw.google_adwords_customer_nme,
			gaw.google_adwords_client_customer_id
		FROM
			"GOOGLE_ADWORDS" as gaw
		INNER JOIN "GOOGLE_ACCOUNT" as gac ON gaw.google_account_id = gac.google_account_id
		WHERE 
			gaw.google_adwords_client_customer_id IN (?);
	`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return googleAccounts, errors.Wrap(err, "sqlx.In error")
	}
	query = googleAccountRepo.Db.Rebind(query)

	//err = googleAccountRepo.Db.Select(&googleAccounts, query, args...)
	rows, err := googleAccountRepo.Db.Query(query, args)

	defer rows.Close()
	for rows.Next() {

	}

	if err != nil {
		return googleAccounts, errors.Wrap(err, "googleAccountRepo.Db.Select error")
	}

	fmt.Printf("%#v", googleAccounts)
	return googleAccounts, nil
}

func (googleAccountRepo *GoogleAccountRepository) FindByCampaignRanByZOWIForZELO() ([]models.GoogleAccount, error) {
	query := `
		select
			distinct on (ga.google_account_id)
			ga.* 
		from "CAMPAIGN" as c
		inner join "LOCATION" as loc on loc.location_id = c.location_id
		inner join "ORGANIZATION" as org on loc.organization_id = org.organization_id
		inner join "GOOGLE_ACCOUNT" as ga on ga.organization_id = org.parent_id
		where c.campaign_ran_by_zowi = true;
	`
	var googleAccouts []models.GoogleAccount
	err := googleAccountRepo.Db.Select(&googleAccouts, query)
	if err != nil {
		return googleAccouts, errors.Wrap(err, "googleAccountRepo.Db.Select error")
	}
	return googleAccouts, nil
}