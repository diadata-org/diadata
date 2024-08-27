package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type Customer struct {
	CustomerID          int       `json:"customer_id"`
	Email               string    `json:"email"`
	AccountCreationDate time.Time `json:"account_creation_date"`
	CustomerPlan        int       `json:"customer_plan"`
	DeployedOracles     int       `json:"deployed_oracles"`
	PaymentStatus       string    `json:"payment_status"`
	// LastPayment         time.Time `json:"last_payment"`
	PaymentSource     string      `json:"payment_source"`
	NumberOfDataFeeds int         `json:"number_of_data_feeds"`
	Active            bool        `json:"active"`
	PublicKeys        []PublicKey `json:"public_keys"`
}

type Plan struct {
	PlanID      int    `json:"plan_id"`
	Name        string `json:"plan_name"`
	Description string `json:"plan_description"`
	Price       int    `json:"plan_price"`
	TotalFeeds  int    `json:"total_feeds"`
}

type PublicKey struct {
	AccessLevel string `json:"access_level"`
	PublicKey   string `json:"public_key"`
	UserName    string `json:"username"`
}

func (rdb *RelDB) SetKeyPair(publickey string, privatekey string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(publickey,privatekey) VALUES ($1,$2) 
	 on conflict(publickey)  
	do
	update set publickey=EXCLUDED.publickey`, keypairTable)
	exec, err := rdb.postgresClient.Exec(context.Background(), query, publickey, privatekey)

	log.Infoln("exec", exec)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetKeyPairID(publicKey string) string {
	query := fmt.Sprintf(`SELECT id from %s WHERE publickey=$1`, keypairTable)
	rows := rdb.postgresClient.QueryRow(context.Background(), query, publicKey)
	var keypairId string

	err := rows.Scan(&keypairId)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}

	return keypairId
}

func (rdb *RelDB) SetOracleConfig(ctx context.Context, customerId, address, feederID, owner, feederAddress, symbols, feedSelection, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, mandatoryFrequency string) error {
	currentTime := time.Now()
	query := fmt.Sprintf(`
		INSERT INTO %s (customer_id, address,feeder_id,owner,symbols,chainID,frequency,sleepseconds,deviationpermille,blockchainnode,mandatory_frequency,feeder_address,createddate,lastupdate,feedSelection) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
		ON CONFLICT(feeder_id)  
		DO UPDATE SET symbols=$4,frequency=$6,sleepseconds=$7,deviationpermille=$8,blockchainnode=$9,mandatory_frequency=$10,feeder_address=$11,lastupdate=$13,feedSelection=$14`,
		oracleconfigTable,
	)

	log.Infoln("SetOracleConfig Query", query)
	_, err := rdb.postgresClient.Exec(ctx, query, customerId, address, feederID, owner, symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, mandatoryFrequency, feederAddress, currentTime, currentTime, feedSelection)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetFeederID(address string) (feederId string) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE owner=$1`, feederaccessTable)
	log.Infoln("GetFeederID query", query)
	log.Infoln("address", address)
	var feederidint int
	err := rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&feederidint)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	feederId = strconv.Itoa(feederidint)

	return
}

func (rdb *RelDB) SetFeederConfig(feederid, oracleconfigid string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (id, oracleconfig_id) 
		VALUES ($1,$2) 
		ON CONFLICT(id)  
		DO UPDATE SET oracleconfig_id=EXCLUDED.oracleconfig_id`,
		feederconfigTable,
	)
	_, err := rdb.postgresClient.Exec(context.Background(), query, feederid, oracleconfigid)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetFeederAccessByID(id string) (owner string) {
	query := fmt.Sprintf(`SELECT owner FROM %s WHERE feeder_id=$1`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&owner)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) GetFeederByID(id string) (owner string) {
	query := fmt.Sprintf(`SELECT owner from %s WHERE feeder_id=$1`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&owner)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}
func (rdb *RelDB) GetFeederLimit(owner string) (limit int) {
	query := fmt.Sprintf(`SELECT total FROM %s WHERE owner=$1`, feederResourceTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, owner).Scan(&limit)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) GetTotalOracles(customerId string) (total int) {
	query := fmt.Sprintf(`SELECT count(*) FROM %s WHERE customer_id=$1 AND active=true`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, customerId).Scan(&total)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) UpdateFeederAddressCheckSum(oracleaddress string) (err error) {

	query := fmt.Sprintf(`
	UPDATE %s 
	SET address=$1
	WHERE address=$2`, oracleconfigTable)

	_, err = rdb.postgresClient.Exec(context.Background(), query, common.HexToAddress(oracleaddress).Hex(), oracleaddress)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) GetExpiredFeeders() (oracleconfigs []dia.OracleConfig, err error) {

	var (
		rows           pgx.Rows
		deviationFloat float64
		query          string
	)

	query = fmt.Sprintf(`
	SELECT 
    t1.address,  t1.feeder_id, t1.deleted, t1.owner, t1.symbols, t1.chainID, 
    t1.frequency, t1.sleepseconds, t1.deviationpermille, t1.blockchainnode, t1.active, 
    t1.mandatory_frequency, t1.feeder_address, t1.createddate, 
    COALESCE(t1.lastupdate, '0001-01-01 00:00:00'::timestamp) AS lastupdate, 
    t1.expired, t1.expired_time,
    COALESCE(MAX(fu.update_time), '0001-01-01 00:00:00'::timestamp) AS max_update_time
FROM %s AS t1
LEFT JOIN %s AS fu 
ON t1.address = fu.oracle_address
GROUP BY  
    t1.address, t1.feeder_id, t1.deleted, t1.owner, t1.symbols, t1.chainID, 
    t1.frequency, t1.sleepseconds, t1.deviationpermille, t1.blockchainnode, t1.active, 
    t1.mandatory_frequency, t1.feeder_address, t1.createddate, 
    t1.lastupdate, t1.expired, t1.expired_time
HAVING 
    EXTRACT(EPOCH FROM (NOW() - lastupdate)) / 86400 > 60
    AND t1.deleted = false
    AND t1.expired = false
	`, oracleconfigTable, feederupdatesTable)

	rows, err = rdb.postgresClient.Query(context.Background(), query)

	if err != nil {
		fmt.Println(err)

		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig     dia.OracleConfig
			symbols          string
			frequencynull    sql.NullString
			sleepsecondsnull sql.NullString
			feedSelection    sql.NullString
		)
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Deleted, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID,
			&frequencynull, &sleepsecondsnull, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode, &oracleconfig.Active,
			&oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress, &oracleconfig.CreatedDate,
			&oracleconfig.LastUpdate, &oracleconfig.Expired, &oracleconfig.ExpiredDate, &oracleconfig.LastOracleUpdate)
		if err != nil {

			log.Error("GetExpiredFeeders scan", err, oracleconfig.FeederID)
		}

		if frequencynull.Valid {
			oracleconfig.Frequency = frequencynull.String

		}

		if feedSelection.Valid {
			oracleconfig.FeederSelection = feedSelection.String
		}

		if sleepsecondsnull.Valid {
			oracleconfig.SleepSeconds = sleepsecondsnull.String

		}

		oracleconfig.Symbols = strings.Split(symbols, ",")
		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return

}
func (rdb *RelDB) GetFeeder(feederID string) (oracleconfig dia.OracleConfig, err error) {
	var (
		row            pgx.Row
		deviationFloat float64
		query          string
	)

	query = fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted,feedselection,expired,expired_time
	FROM %s  WHERE feeder_id=$1
	`, oracleconfigTable)
	row = rdb.postgresClient.QueryRow(context.Background(), query, feederID)

	if err != nil {
		return
	}

	var (
		symbols          string
		frequencynull    sql.NullString
		sleepsecondsnull sql.NullString
		feedSelection    sql.NullString
	)
	err = row.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &frequencynull, &sleepsecondsnull, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress, &oracleconfig.CreatedDate, &oracleconfig.LastUpdate, &oracleconfig.Deleted, &feedSelection, &oracleconfig.Expired, &oracleconfig.ExpiredDate)
	if err != nil {

		log.Error("GetFeeder scan", err, oracleconfig.FeederID)
	}

	if frequencynull.Valid {
		oracleconfig.Frequency = frequencynull.String

	}

	if feedSelection.Valid {
		oracleconfig.FeederSelection = feedSelection.String
	}

	if sleepsecondsnull.Valid {
		oracleconfig.SleepSeconds = sleepsecondsnull.String

	}

	oracleconfig.Symbols = strings.Split(symbols, ",")
	if oracleconfig.DeviationPermille != "" {
		deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
		if err != nil {
			log.Error(err)

		}
		oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
	}

	return
}

func (rdb *RelDB) GetAllFeeders(isDeleted bool, isExpired bool) (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows           pgx.Rows
		deviationFloat float64
		query          string
	)

	switch {

	case isDeleted:
		{
			query = fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted,feedselection,expired,expired_time
	FROM %s  WHERE mandatory_frequency  IS NOT NULL   and deleted=$1
	`, oracleconfigTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query, isDeleted)

		}

	case isExpired:
		{
			query = fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted,feedselection,expired,expired_time
	FROM %s  WHERE mandatory_frequency  IS NOT NULL   and expired=$1
	`, oracleconfigTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query, isExpired)
		}

	case isExpired && isDeleted:
		{
			query = fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted,feedselection,expired,expired_time
	FROM %s  WHERE mandatory_frequency  IS NOT NULL   and expired=$1 and deleted=$2
	`, oracleconfigTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query, isExpired, isDeleted)
		}
	case !isExpired && !isDeleted:
		{
			query = fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted,feedselection,expired,expired_time
	FROM %s  WHERE mandatory_frequency  IS NOT NULL 
	`, oracleconfigTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query)

		}

	}

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig     dia.OracleConfig
			symbols          string
			frequencynull    sql.NullString
			sleepsecondsnull sql.NullString
			feedSelection    sql.NullString
		)
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &frequencynull, &sleepsecondsnull, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress, &oracleconfig.CreatedDate, &oracleconfig.LastUpdate, &oracleconfig.Deleted, &feedSelection, &oracleconfig.Expired, &oracleconfig.ExpiredDate)
		if err != nil {

			log.Error("GetAllFeeders scan", err, oracleconfig.FeederID)
		}

		if frequencynull.Valid {
			oracleconfig.Frequency = frequencynull.String

		}

		if feedSelection.Valid {
			oracleconfig.FeederSelection = feedSelection.String
		}

		if sleepsecondsnull.Valid {
			oracleconfig.SleepSeconds = sleepsecondsnull.String

		}

		oracleconfig.Symbols = strings.Split(symbols, ",")
		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}
func (rdb *RelDB) GetFeederResources() (addresses []string, err error) {
	var (
		rows pgx.Rows
	)
	query := fmt.Sprintf(`
	SELECT owner
	FROM %s`, feederResourceTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			address string
		)
		err := rows.Scan(&address)
		if err != nil {
			log.Error(err)
		}
		addresses = append(addresses, address)
	}
	return

}

func (rdb *RelDB) GetOraclesByCustomer(customerId string) (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows           pgx.Rows
		deviationFloat float64
	)

	/*
	 `SELECT address, chainid,  COALESCE(latest.scraped_block, 0) AS latest_scraped_block FROM oracleconfig
	  LEFT JOIN (SELECT oracle_address, chain_id, MAX(update_block) AS scraped_block FROM feederupdates GROUP BY oracle_address,chain_id) latest ON (oracleconfig.address = latest.oracle_address and oracleconfig.chainid = latest.chain_id) WHERE  oracleconfig.chainid = '%s'`


	*/

	query := fmt.Sprintf(`
	SELECT 
    	t1.address,  t1.feeder_id,t1.deleted,  t1.owner,  t1.symbols,  t1.chainID, 
    	t1.frequency,  t1.sleepseconds,  t1.deviationpermille,  t1.blockchainnode,   t1.active, 
    	t1.mandatory_frequency,  t1.feeder_address,  t1.createddate,  t1.feedselection, 
    	COALESCE(t1.lastupdate, '0001-01-01 00:00:00'::timestamp) AS lastupdate, 
		t1.expired, t1.expired_time,
    	COALESCE(MAX(fu.update_time), '0001-01-01 00:00:00'::timestamp) AS max_update_time
	FROM %s AS t1
	LEFT JOIN %s AS fu 
    ON t1.address = fu.oracle_address 
    AND t1.chainID = fu.chain_id
	WHERE t1.customer_id = $1 
	GROUP BY  
		t1.address,  t1.feeder_id, t1.deleted, t1.owner,  t1.symbols,  t1.chainID, 
   		t1.frequency,  t1.sleepseconds,  t1.deviationpermille,  t1.blockchainnode,  t1.active, 
		t1.mandatory_frequency,  t1.feeder_address, t1.createddate, t1.feedselection, 
     	t1.lastupdate, t1.expired,t1.expired_time;`, oracleconfigTable, feederupdatesTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, customerId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig  dia.OracleConfig
			symbols       string
			feedSelection sql.NullString
		)

		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Deleted, &oracleconfig.Owner, &symbols,
			&oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille,
			&oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress,
			&oracleconfig.CreatedDate, &feedSelection, &oracleconfig.LastUpdate, &oracleconfig.Expired, &oracleconfig.ExpiredDate, &oracleconfig.LastOracleUpdate)
		if err != nil {
			log.Error(err)
		}
		if feedSelection.Valid {
			oracleconfig.FeederSelection = feedSelection.String
		}

		oracleconfig.Symbols = strings.Split(symbols, ",")

		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		lastupdate := oracleconfig.LastOracleUpdate

		if oracleconfig.LastOracleUpdate.IsZero() {
			lastupdate = oracleconfig.CreatedDate
		}
		oracleconfig.ExpiringDate = lastupdate.Add(time.Duration(60 * time.Hour * 24))
		if oracleconfig.ExpiringDate.Before(time.Now()) {
			oracleconfig.Expired = true
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}

func (rdb *RelDB) GetOraclesByOwner(owner string) (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows           pgx.Rows
		deviationFloat float64
	)

	/*
	 `SELECT address, chainid,  COALESCE(latest.scraped_block, 0) AS latest_scraped_block FROM oracleconfig
	  LEFT JOIN (SELECT oracle_address, chain_id, MAX(update_block) AS scraped_block FROM feederupdates GROUP BY oracle_address,chain_id) latest ON (oracleconfig.address = latest.oracle_address and oracleconfig.chainid = latest.chain_id) WHERE  oracleconfig.chainid = '%s'`


	*/

	query := fmt.Sprintf(`
	SELECT 
    	t1.address,  t1.feeder_id,t1.deleted,  t1.owner,  t1.symbols,  t1.chainID, 
    	t1.frequency,  t1.sleepseconds,  t1.deviationpermille,  t1.blockchainnode,   t1.active, 
    	t1.mandatory_frequency,  t1.feeder_address,  t1.createddate,  t1.feedselection, 
    	COALESCE(t1.lastupdate, '0001-01-01 00:00:00'::timestamp) AS lastupdate, 
		t1.expired, t1.expired_time,
    	COALESCE(MAX(fu.update_time), '0001-01-01 00:00:00'::timestamp) AS max_update_time
	FROM %s AS t1
	LEFT JOIN %s AS fu 
    ON t1.address = fu.oracle_address 
    AND t1.chainID = fu.chain_id
	WHERE t1.owner = $1 
	GROUP BY  
		t1.address,  t1.feeder_id, t1.deleted, t1.owner,  t1.symbols,  t1.chainID, 
   		t1.frequency,  t1.sleepseconds,  t1.deviationpermille,  t1.blockchainnode,  t1.active, 
		t1.mandatory_frequency,  t1.feeder_address, t1.createddate, t1.feedselection, 
     	t1.lastupdate, t1.expired,t1.expired_time;`, oracleconfigTable, feederupdatesTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, owner)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig  dia.OracleConfig
			symbols       string
			feedSelection sql.NullString
		)

		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Deleted, &oracleconfig.Owner, &symbols,
			&oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille,
			&oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress,
			&oracleconfig.CreatedDate, &feedSelection, &oracleconfig.LastUpdate, &oracleconfig.Expired, &oracleconfig.ExpiredDate, &oracleconfig.LastOracleUpdate)
		if err != nil {
			log.Error(err)
		}
		if feedSelection.Valid {
			oracleconfig.FeederSelection = feedSelection.String
		}

		oracleconfig.Symbols = strings.Split(symbols, ",")

		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		lastupdate := oracleconfig.LastOracleUpdate

		if oracleconfig.LastOracleUpdate.IsZero() {
			lastupdate = oracleconfig.CreatedDate
		}
		oracleconfig.ExpiringDate = lastupdate.Add(time.Duration(60 * time.Hour * 24))
		if oracleconfig.ExpiringDate.Before(time.Now()) {
			oracleconfig.Expired = true
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}

func (rdb *RelDB) GetOracleConfig(address, chainid string) (oracleconfig dia.OracleConfig, err error) {
	var (
		symbols       string
		feedSelection sql.NullString
	)
	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,feeder_address,symbols, chainid, feedSelection,deviationpermille, sleepseconds,frequency, blockchainnode, mandatory_frequency
	FROM %s 
	WHERE address=$1 and chainid=$2`, oracleconfigTable)
	fmt.Println(query)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, chainid).Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &oracleconfig.FeederAddress, &symbols, &oracleconfig.ChainID, &feedSelection, &oracleconfig.DeviationPermille, &oracleconfig.SleepSeconds, &oracleconfig.Frequency, &oracleconfig.BlockchainNode, &oracleconfig.MandatoryFrequency)
	if err != nil {
		return
	}

	if feedSelection.Valid {
		oracleconfig.FeederSelection = feedSelection.String
	}

	oracleconfig.Symbols = strings.Split(symbols, " ")

	return
}

func (rdb *RelDB) ChangeOracleState(feederID string, active bool) (err error) {
	currentTime := time.Now()

	query := fmt.Sprintf(`
	UPDATE %s 
	SET active=$1, lastupdate=$3
	WHERE feeder_id=$2`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, active, feederID, currentTime)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) DeleteOracle(feederID string) (err error) {
	currentTime := time.Now()
	query := fmt.Sprintf(`
	UPDATE %s 
	SET deleted=$1,lastupdate=$3
	WHERE feeder_id=$2`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, true, feederID, currentTime)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) ExpireOracle(feederID string) (err error) {
	currentTime := time.Now()
	query := fmt.Sprintf(`
	UPDATE %s 
	SET expired=$1, deleted=$2,lastupdate=$4
	WHERE feeder_id=$3`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, true, true, feederID, currentTime)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) GetOracleUpdateCount(address, chainid, symbol string) (int64, error) {
	query := fmt.Sprintf(`
	SELECT  count(*) from %s
	WHERE oracle_address=$1 and chain_id=$2
	AND ($3 = '' OR asset_key = $3)
	
	`, feederupdatesTable)

	var numUpdates sql.NullInt64
	err := rdb.postgresClient.QueryRow(context.Background(), query, address, chainid, symbol).Scan(&numUpdates)
	if numUpdates.Valid {
		return numUpdates.Int64, nil
	}
	return 0, err

}

func (rdb *RelDB) GetOracleLastUpdate(address, chainid, symbol string) (time.Time, string, error) {
	symbol = symbol + "/USD"
	var (
		updateTime sql.NullTime
		price      string
	)
	query := fmt.Sprintf(`
	SELECT update_time,asset_price
	FROM %s fu
	WHERE oracle_address = $1 AND chain_id = $2 and asset_key=$3 order by update_block desc LIMIT 1 
	`, feederupdatesTable)

	err := rdb.postgresClient.QueryRow(context.Background(), query, address, chainid, symbol).Scan(&updateTime, &price)
	if err != nil {
		return time.Time{}, "", err
	}

	return updateTime.Time, price, nil
}
func (rdb *RelDB) GetOracleUpdates(address string, chainid string, offset int) ([]dia.OracleUpdate, error) {
	query := fmt.Sprintf(`
	SELECT fu.oracle_address,
		fu.transaction_hash,
		fu.transaction_cost,
		fu.asset_key,
		fu.asset_price,
		fu.update_block,
		fu.update_from,
		fu.from_balance,
		fu.gas_cost,
		fu.gas_used,
		fu.chain_id,
		fu.update_time,
		oc.creation_block,
		oc.creation_block_time


	FROM %s fu
	JOIN %s oc ON fu.oracle_address = oc.address AND fu.chain_id = oc.chainID 
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2 order by fu.update_block desc LIMIT 20 OFFSET %d
	`, feederupdatesTable, oracleconfigTable, offset)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		updates []dia.OracleUpdate
	)

	for rows.Next() {

		var (
			update            dia.OracleUpdate
			updateTime        sql.NullTime
			creationBlock     sql.NullInt64
			creationBlockTime sql.NullTime
		)
		err := rows.Scan(
			&update.OracleAddress,
			&update.TransactionHash,
			&update.TransactionCost,
			&update.AssetKey,
			&update.AssetPrice,
			&update.UpdateBlock,
			&update.UpdateFrom,
			&update.FromBalance,
			&update.GasCost,
			&update.GasUsed,
			&update.ChainID,
			&updateTime,
			&creationBlock,
			&creationBlockTime,
		)

		if updateTime.Valid {
			update.UpdateTime = updateTime.Time
		}
		if creationBlockTime.Valid {
			update.CreationBlockTime = creationBlockTime.Time
		}
		if creationBlock.Valid {
			update.CreationBlock = uint64(creationBlock.Int64)
		}

		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return updates, nil
}
func (rdb *RelDB) GetOracleUpdatesByTimeRange(address, chainid, symbol string, offset int, startTime, endTime time.Time) ([]dia.OracleUpdate, error) {
	query := ""
	if offset == -1 {
		query = fmt.Sprintf(`
	SELECT fu.oracle_address,
		fu.transaction_hash,
		fu.transaction_cost,
		fu.asset_key,
		fu.asset_price,
		fu.update_block,
		fu.update_from,
		fu.from_balance,
		fu.gas_cost,
		fu.gas_used,
		fu.chain_id,
		fu.update_time,
		oc.creation_block,
		oc.creation_block_time

	FROM %s fu
	JOIN %s oc ON fu.oracle_address = oc.address AND fu.chain_id = oc.chainID 
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2 
    AND ($3 = 0 OR $4 = 0 OR (fu.update_time < to_timestamp($3) AND fu.update_time > to_timestamp($4)))
	AND ($5 = '' OR fu.asset_key = $5)  

	order by fu.update_block desc 
	`, feederupdatesTable, oracleconfigTable)

	} else {
		query = fmt.Sprintf(`
	SELECT fu.oracle_address,
		fu.transaction_hash,
		fu.transaction_cost,
		fu.asset_key,
		fu.asset_price,
		fu.update_block,
		fu.update_from,
		fu.from_balance,
		fu.gas_cost,
		fu.gas_used,
		fu.chain_id,
		fu.update_time,
		oc.creation_block,
		oc.creation_block_time

	FROM %s fu
	JOIN %s oc ON fu.oracle_address = oc.address AND fu.chain_id = oc.chainID 
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2 
    AND ($3 = 0 OR $4 = 0 OR (fu.update_time < to_timestamp($3) AND fu.update_time > to_timestamp($4)))
	AND ($5 = '' OR fu.asset_key = $5)  

	order by fu.update_block desc LIMIT 20 OFFSET %d
	`, feederupdatesTable, oracleconfigTable, offset)

	}

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid, endTime.Unix(), startTime.Unix(), symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		updates []dia.OracleUpdate
	)

	for rows.Next() {

		var (
			update            dia.OracleUpdate
			updateTime        sql.NullTime
			creationBlock     sql.NullInt64
			creationBlockTime sql.NullTime
		)
		err := rows.Scan(
			&update.OracleAddress,
			&update.TransactionHash,
			&update.TransactionCost,
			&update.AssetKey,
			&update.AssetPrice,
			&update.UpdateBlock,
			&update.UpdateFrom,
			&update.FromBalance,
			&update.GasCost,
			&update.GasUsed,
			&update.ChainID,
			&updateTime,
			&creationBlock,
			&creationBlockTime,
		)

		if updateTime.Valid {
			update.UpdateTime = updateTime.Time
		}
		if creationBlockTime.Valid {
			update.CreationBlockTime = creationBlockTime.Time
		}
		if creationBlock.Valid {
			update.CreationBlock = uint64(creationBlock.Int64)
		}

		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return updates, nil
}
func (rdb *RelDB) GetLastOracleUpdate(address string, chainid string) ([]dia.OracleUpdate, error) {
	query := fmt.Sprintf(`
	SELECT fu.oracle_address,
		fu.transaction_hash,
		fu.transaction_cost,
		fu.asset_key,
		fu.asset_price,
		fu.update_block,
		fu.update_from,
		fu.from_balance,
		fu.gas_cost,
		fu.gas_used,
		fu.chain_id,
		fu.update_time,
		oc.creation_block,
		oc.creation_block_time


	FROM %s fu
	JOIN %s oc ON fu.oracle_address = oc.address AND fu.chain_id = oc.chainID 
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2 order by fu.update_time desc LIMIT 1
	`, feederupdatesTable, oracleconfigTable)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		updates []dia.OracleUpdate
	)

	for rows.Next() {

		var (
			update            dia.OracleUpdate
			updateTime        sql.NullTime
			creationBlock     sql.NullInt64
			creationBlockTime sql.NullTime
		)
		err := rows.Scan(
			&update.OracleAddress,
			&update.TransactionHash,
			&update.TransactionCost,
			&update.AssetKey,
			&update.AssetPrice,
			&update.UpdateBlock,
			&update.UpdateFrom,
			&update.FromBalance,
			&update.GasCost,
			&update.GasUsed,
			&update.ChainID,
			&updateTime,
			&creationBlock,
			&creationBlockTime,
		)

		if updateTime.Valid {
			update.UpdateTime = updateTime.Time
		}
		if creationBlockTime.Valid {
			update.CreationBlockTime = creationBlockTime.Time
		}
		if creationBlock.Valid {
			update.CreationBlock = uint64(creationBlock.Int64)
		}

		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return updates, nil
}
func (rdb *RelDB) GetTotalGasSpend(address string, chainid string) (float64, error) {
	query := fmt.Sprintf(`
	SELECT 
		SUM(fu.gas_used) AS total_gas_used
	FROM %s fu
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2
	`, feederupdatesTable)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		fmt.Println("err", err)
		return 0.0, err
	}
	defer rows.Close()

	var (
		gasUsed sql.NullFloat64
	)

	for rows.Next() {

		rows.Scan(

			&gasUsed,
		)

	}

	if gasUsed.Valid {
		return gasUsed.Float64, nil

	}
	return 0.0, nil

}

func (rdb *RelDB) GetBalanceRemaining(address string, chainid string) (float64, error) {
	query := fmt.Sprintf(`
	SELECT 
		from_balance,
		update_time
	FROM %s fu
	WHERE fu.oracle_address = $1 AND fu.chain_id = $2 ORDER BY update_time DESC LIMIT 1
	`, feederupdatesTable)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		fmt.Println("err", err)
		return 0.0, err
	}
	defer rows.Close()

	var (
		gasRemaining sql.NullFloat64
	)

	for rows.Next() {

		rows.Scan(
			&gasRemaining,
			nil,
		)

	}

	if gasRemaining.Valid {
		return gasRemaining.Float64, nil

	}
	return 0.0, nil

}

func (rdb *RelDB) GetDayWiseUpdates(address string, chainid string) ([]dia.FeedUpdates, float64, float64, error) {
	query := fmt.Sprintf(`
	WITH Dates AS (
		SELECT generate_series(CURRENT_DATE - INTERVAL '29 days', CURRENT_DATE, INTERVAL '1 day')::DATE AS day
	),
	DailyUpdates AS (
		SELECT 
			d.day, 
			AVG(fu.gas_used) AS average_gas_used,
			SUM(fu.gas_used) AS total_gas_used,
			COUNT(fu.update_time) AS total_updates
		FROM Dates d
		LEFT JOIN %s fu ON DATE(fu.update_time) = d.day
			AND fu.oracle_address = $1 
			AND fu.chain_id = $2
		GROUP BY d.day
	)
	SELECT 
		day, 
		COALESCE(total_gas_used, 0) AS total_gas_used,
		COALESCE(total_updates, 0) AS total_updates,
		(SELECT AVG(total_updates::NUMERIC) FROM DailyUpdates) AS average_total_updates_per_day,
		(SELECT AVG(average_gas_used) FROM DailyUpdates) AS average_gas_used_per_day
	FROM DailyUpdates
	ORDER BY day DESC;
	`, feederupdatesTable)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		fmt.Println("err", err)
		return nil, 0, 0, err
	}
	defer rows.Close()

	var (
		avgGasUsed     sql.NullFloat64
		avgUpdateCount sql.NullFloat64
		stats          []dia.FeedUpdates
		updateTime     sql.NullTime
	)

	for rows.Next() {
		du := dia.FeedUpdates{}
		rows.Scan(
			&updateTime,
			&du.GasUsed,
			&du.UpdateCount,
			&avgUpdateCount,
			&avgGasUsed,
		)
		if updateTime.Valid {
			du.Day = updateTime.Time
		}

		stats = append(stats, du)

	}

	return stats, avgGasUsed.Float64, avgUpdateCount.Float64, nil

}

// dave oracleconfig in cache
func (datastore *DB) SetOracleConfigCache(oc dia.OracleConfig) error {
	key := oc.Address + "-" + oc.ChainID

	if datastore.redisClient == nil {
		return nil
	}

	data, _ := json.Marshal(oc)

	err := datastore.redisClient.Set(key, data, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetOracleConfigCache %v\n", err, oc)
	}

	return err
}

func (datastore *DB) GetOracleConfigCache(key string) (dia.OracleConfig, error) {
	// key := oc.Address + oc.ChainID
	var oc dia.OracleConfig
	if datastore.redisClient == nil {
		return oc, errors.New("no redisclient")
	}

	err := datastore.redisClient.Get(key).Scan(&oc)
	if err != nil {
		log.Printf("Error: %v on GetOracle--ConfigCache %v\n", err, oc)
		return oc, errors.New("no redisclient")

	}

	return oc, err
}

func (reldb *RelDB) AddWalletKeys(owner, username, accessLevel string, publicKey []string) error {
	var err error

	customerID, err := reldb.GetCustomerIDByWalletPublicKey(owner)
	if err != nil {
		return err
	}

	tx, err := reldb.postgresClient.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err

	}
	defer tx.Rollback(context.Background())

	for _, publicKey := range publicKey {
		_, err = tx.Exec(context.Background(), `
			INSERT INTO wallet_public_keys (customer_id, public_key, access_level,username)
			VALUES ($1, $2,$3,$4)
		`, customerID, publicKey, accessLevel, username)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (reldb *RelDB) UpdateAccessLevel(username, accessLevel, publicKey string) error {
	var err error

	tx, err := reldb.postgresClient.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err

	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `
		UPDATE wallet_public_keys
		SET access_level = $1,
			    username = $3
		WHERE public_key = $2
		AND ($1 IS NOT NULL AND $1 != '' OR $3 IS NOT NULL AND $3 != '');

		`, accessLevel, publicKey, username)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (reldb *RelDB) RemoveWalletKeys(publicKey []string) error {
	var err error

	tx, err := reldb.postgresClient.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err

	}
	defer tx.Rollback(context.Background())

	for _, publicKey := range publicKey {
		_, err = tx.Exec(context.Background(), `
		DELETE FROM wallet_public_keys 
		WHERE public_key = $1
		`, publicKey)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
func (reldb *RelDB) GetCustomerIDByWalletPublicKey(publicKey string) (int, error) {
	var customerID int
	err := reldb.postgresClient.QueryRow(context.Background(), `
		SELECT customer_id
		FROM wallet_public_keys
		WHERE public_key = $1
	`, publicKey).Scan(&customerID)
	if err != nil {
		return 0, err
	}
	return customerID, nil
}

func (reldb *RelDB) CreateCustomer(email string, customerPlan int, paymentStatus string, paymentSource string, numberOfDataFeeds int, walletPublicKeys []string) error {
	tx, err := reldb.postgresClient.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("unable to start a transaction: %v", err)
	}
	defer tx.Rollback(context.Background())

	// Insert the new customer
	var customerID int
	insertCustomerQuery := `
        INSERT INTO customers (email, customer_plan, payment_status, payment_source, number_of_data_feeds)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING customer_id`

	err = tx.QueryRow(context.Background(), insertCustomerQuery, email, customerPlan, paymentStatus, paymentSource, numberOfDataFeeds).Scan(&customerID)
	if err != nil {
		return fmt.Errorf("unable to insert customer: %v", err)
	}

	err = addWalletPublicKeys(tx, customerID, walletPublicKeys)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("unable to commit transaction: %v", err)
	}

	return nil
}

func (reldb *RelDB) UpdateCustomerPlan(ctx context.Context, customerID int, customerPlan int, paymentSource string, lastPayment string) error {

	ut, err := strconv.ParseInt(lastPayment, 10, 64)
	lastPaymentts := time.Unix(ut, 0)

	query := `
        UPDATE customers
        SET customer_plan = $1,
            payment_source = $2,
            last_payment = $3
        WHERE customer_id = $4
    `
	_, err = reldb.postgresClient.Exec(ctx, query, customerPlan, paymentSource, lastPaymentts, customerID)
	return err
}

func addWalletPublicKeys(tx pgx.Tx, customerID int, walletPublicKeys []string) error {
	insertWalletKeyQuery := `
        INSERT INTO wallet_public_keys (customer_id, public_key)
        VALUES ($1, $2)`

	for _, publicKey := range walletPublicKeys {
		_, err := tx.Exec(context.Background(), insertWalletKeyQuery, customerID, publicKey)
		if err != nil {
			return fmt.Errorf("unable to insert wallet public key: %v", err)
		}
	}

	return nil
}

func (reldb *RelDB) GetPlan(ctx context.Context, planID int) (*Plan, error) {
	var plan Plan
	query := `SELECT plan_id, plan_name, plan_description, plan_price, total_feeds
	       FROM plans
		   WHERE plan_id = $1`
	err := reldb.postgresClient.QueryRow(ctx, query, planID).Scan(
		&plan.PlanID,
		&plan.Name,
		&plan.Description,
		&plan.Price,
		&plan.TotalFeeds,
	)
	if err != nil {
		return nil, err
	}
	return &plan, err

}

func (reldb *RelDB) GetCustomerByPublicKey(publicKey string) (*Customer, error) {
	var customer Customer
	query := `
		SELECT c.customer_id, c.email, c.account_creation_date, c.customer_plan, c.deployed_oracles,
		       c.payment_status, c.payment_source, c.number_of_data_feeds, c.active
		FROM customers c
		INNER JOIN wallet_public_keys wpk ON c.customer_id = wpk.customer_id
		WHERE wpk.public_key = $1
	`
	var (
		emailSql         sql.NullString
		customerPlanSql  sql.NullInt16
		paymentStatusSql sql.NullString
		paymentSourceSql sql.NullString
	)
	err := reldb.postgresClient.QueryRow(context.Background(), query, publicKey).Scan(
		&customer.CustomerID,
		&emailSql,
		&customer.AccountCreationDate,
		&customerPlanSql,
		&customer.DeployedOracles,
		&paymentStatusSql,
		// &customer.LastPayment,
		&paymentSourceSql,
		&customer.NumberOfDataFeeds,
		&customer.Active,
	)
	if err != nil {
		return nil, err
	}
	query = `
		SELECT public_key, access_level,username
		FROM wallet_public_keys
		WHERE customer_id = $1
	`
	rows, err := reldb.postgresClient.Query(context.Background(), query, customer.CustomerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var username sql.NullString

	for rows.Next() {
		var publicKey PublicKey
		if err := rows.Scan(&publicKey.PublicKey, &publicKey.AccessLevel, &username); err != nil {
			return nil, err
		}
		if username.Valid {
			publicKey.UserName = username.String
		}
		customer.PublicKeys = append(customer.PublicKeys, publicKey)
	}

	if emailSql.Valid {
		customer.Email = emailSql.String
	}
	if paymentStatusSql.Valid {
		customer.PaymentStatus = paymentStatusSql.String
	}
	if customerPlanSql.Valid {
		customer.CustomerPlan = int(customerPlanSql.Int16)
	}
	if paymentSourceSql.Valid {
		customer.PaymentSource = paymentSourceSql.String
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &customer, nil
}

func (reldb *RelDB) GetAccessLevel(publicKey string) (string, error) {
	var accessLevel string
	err := reldb.postgresClient.QueryRow(context.Background(), `
		SELECT access_level
		FROM wallet_public_keys
		WHERE public_key = $1
	`, publicKey).Scan(&accessLevel)

	if err != nil {
		return "", err
	}
	return accessLevel, nil

}
