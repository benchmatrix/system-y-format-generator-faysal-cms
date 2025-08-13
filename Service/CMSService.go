package service

import (
	"database/sql"
	"log"
	"system-y-format-generator-faysal-cms/Models/CMS"
)

func GetNewRMGObservationsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.RMGNewObservationDetails {

	query := `SELECT
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1164') as companyshare,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1194') as natureofrequest,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1195') as deferal,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1165') as collateralevaluation,
		
		(SELECT TOP 1 ISNULL(cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1199') as financialhighlights,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1216') as projectionbase,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY= '1217') as projectionscenario,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1218') as varianceanalysis,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '63') as riskobservation,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '115') as internalindicators,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1169') as returnoncapital,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1170') as internalexternal,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where a.ACTIVITYFIELDSETUP_KEY = '1219') as additionalinformation,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where  a.ACTIVITYFIELDSETUP_KEY = '1203') as recommendation,
		
		(SELECT TOP 1 ISNULL(cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1  
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where  a.ACTIVITYFIELDSETUP_KEY = '1161') as riskraing,
		
		(SELECT TOP 1 ISNULL( cu.ACTIVITYFIELDSETUP_VALUE,'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where  a.ACTIVITYFIELDSETUP_KEY = '1144') as industrystrategy,

		(SELECT TOP 1 ISNULL(ISNULL(cu.ACTIVITYFIELDSETUP_VALUE, cu.ACTIVITYFIELDSETUP_VALUE),'-') from CONFIG_ACTIVITYFIELDSETUP a
		left join CUSTOMER_ACTIVITY_FIELDSAPPROVED cu on cu.ACTIVITYFIELDSETUP_ID =a.ACTIVITYFIELDSETUP_ID AND cu.CUSTOMERCODE = ?1 
		left join CONFIG_STATUS s on s.STATUS_ID= a.STATUS_ID
		where  a.ACTIVITYFIELDSETUP_KEY = '1256') as environmentalrisk
		from config_menu m
		left join CONFIG_STATUS s on s.STATUS_ID= m.STATUS_ID
		where m.menu_key='RMGMemo' AND s.STATUS_ID = 1`

	rows, err := Db.Query(query, referenceId)
	if err != nil {
		log.Printf("Error executing GetNewRMGObservationsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var rmgList []CMS.RMGNewObservationDetails
	for rows.Next() {
		var item CMS.RMGNewObservationDetails
		err = rows.Scan(
			&item.Companyshare,
			&item.Natureofrequest,
			&item.Deferal,
			&item.Collateralevaluation,
			&item.Financialhighlights,
			&item.Projectionbase,
			&item.Projectionscenario,
			&item.Varianceanalysis,
			&item.Riskobservation,
			&item.Internalindicators,
			&item.Returnoncapital,
			&item.Internalexternal,
			&item.Additionalinformation,
			&item.Recommendation,
			&item.RiskRating,
			&item.IndustryStrategy,
			&item.EnviornmentalRisk,
		)
		if err != nil {
			log.Printf("Error scanning GetNewRMGObservationsDetails row: %v", err)
			continue
		}
		rmgList = append(rmgList, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after row iteration in GetNewRMGObservationsDetails: %v", err)
	}

	return rmgList
}

func GetTermsAndConditionsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, alertActivityTypeKey string, referenceId int, fromApproved bool, ReportTable string) []CMS.TermsandConditionsDetails {

	query := `SELECT ISNULL(TAC.TITLE,'-') AS title,ISNULL(OT.TITLE,'-') AS conditiontype,ISNULL(TAC.DESCRIPTION,'-') AS description,ISNULL(ES.TITLE,'-') AS STATUS,ISNULL(OPT.TITLE,'') as complied
		FROM CUSTOMER_TERMSANDCONDITIONS TAC LEFT JOIN OBSERVATION_TYPE OT ON OT.OBSERVATIONTYPE_ID = TAC.OBSERVATIONTYPE_ID 
		LEFT JOIN CONFIG_ENTITYSTATUS ES ON ES.ENTITYSTATUS_ID = TAC.ENTITYSTATUS_ID LEFT JOIN OPTIONS OPT ON OPT.OPTIONS_ID = ISNULL(TAC.COMPLIANTNONCOMPLIANT,2)
		WHERE TAC.PARENT_REFERENCEID = ?1   and TAC.CUSTOMERPOLICYCOMPLIANCE_ID IS NULL 
		UNION ALL		
		SELECT ISNULL(TAC.TITLE,'-') AS title,ISNULL(OT.TITLE,'-') AS conditiontype, ISNULL(TAC.DESCRIPTION,'-') AS description,ISNULL(ES.TITLE,'-') AS STATUS,ISNULL(OPT.TITLE,'') as complied
		FROM CUSTOMER_TERMSANDCONDITIONS TAC LEFT JOIN OBSERVATION_TYPE OT ON OT.OBSERVATIONTYPE_ID = TAC.OBSERVATIONTYPE_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ES ON ES.ENTITYSTATUS_ID = TAC.ENTITYSTATUS_ID
		LEFT JOIN CUSTOMERAPPROVED C ON C.CUSTOMER_ID = TAC.PARENT_REFERENCEID 
		LEFT JOIN CUSTOMER_POLICYCOMPLIANCE CP ON CP.CUSTOMERPOLICYCOMPLIANCE_ID = TAC.CUSTOMERPOLICYCOMPLIANCE_ID
		LEFT JOIN CONFIG_POLICYANDEXCEPTIONSETUP PES ON PES.POLICYANDEXCEPTIONSETUP_ID = CP.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE PERS ON PERS.REGULATORYSEGMENT_ID=C.REGULATORYSEGMENT_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PERS.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_POLICYEXCEPTIONSETUPBUSINESSSEGMENTLINKAGES PEBS ON PEBS.BUSINESSHIERARCHY_ID= C.BUSINESSHIERARCHY_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PEBS.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN OPTIONS OPT ON OPT.OPTIONS_ID = ISNULL(TAC.COMPLIANTNONCOMPLIANT,2) 
		WHERE TAC.PARENT_REFERENCEID = ?1   and TAC.CUSTOMERPOLICYCOMPLIANCE_ID IS NOT NULL
		
		and isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0)= case when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) !=0 and ( PERS.POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE_ID is not null 
		and PEBS.POLICYEXCEPTIONSETUPBUSINESS_ID is not null) then PES.POLICYANDEXCEPTIONSETUP_ID  when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) =0 then 0 end`

	rows, err := Db.Query(query, referenceId)
	if err != nil {
		log.Printf("Error executing GetTermsAndConditionsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var results []CMS.TermsandConditionsDetails
	for rows.Next() {
		var item CMS.TermsandConditionsDetails
		err = rows.Scan(&item.Title, &item.Type, &item.Description, &item.Status, &item.Complied)
		if err != nil {
			log.Printf("Error scanning GetTermsAndConditionsDetails row: %v", err)
			continue
		}
		results = append(results, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTermsAndConditionsDetails: %v", err)
	}

	return results
}

func GetDeferralDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.DeferalWaiversOpenExceptionDetails {

	query := `Select ISNULL(CL.CUSTOMERPOLICYCOMPLIANCE_KEY,CP.CUSTOMERPOLICYCOMPLIANCE_KEY) AS REF  ,  ISNULL(CE.TITLE,CL.CUSTOMERPOLICYCOMPLIANCE_TITLE) as title,
		ISNULL(CONVERT(varchar,ISNULL(CLR.DEFFERAL_DAYS,CL.DEFFERAL_DAYS)),'-') AS EXISTING_DEFFERAL_DAYS ,
		ISNULL(CONVERT(varchar,ISNULL(CLR.DEFERRALCOUNT,CL.DEFERRALCOUNT)),'-') AS EXISTING_DEFERRALCOUNT ,
		FORMAT(ISNULL(CLR.DEFFERAL_UPTO,CL.DEFFERAL_UPTO),'MMMM dd, yyyy') AS EXISTING_DEFERRALEXPIRY ,
		CASE WHEN (CP.IS_COMPLIANT=1 or CS.ENTITYSTATUS_KEY='Waived' or  CS.ENTITYSTATUS_KEY='open') THEN '' ELSE ISNULL(CONVERT(varchar,CP.DEFFERAL_DAYS),'') END AS PROPOSED_DEFFERAL_DAYS ,
		CASE WHEN (CP.IS_COMPLIANT=1 or CS.ENTITYSTATUS_KEY='Waived' or  CS.ENTITYSTATUS_KEY='open') THEN '' ELSE ISNULL(CONVERT(varchar,CP.DEFERRALCOUNT),'')  END AS PROPOSED_DEFERRALCOUNT ,
		CASE WHEN (CP.IS_COMPLIANT=1 or CS.ENTITYSTATUS_KEY='Waived' or  CS.ENTITYSTATUS_KEY='open') THEN '' ELSE FORMAT(CP.DEFFERAL_UPTO,'MMMM dd, yyyy') END AS PROPOSED_DEFERRALEXPIRY ,
		CASE WHEN CP.IS_COMPLIANT=1 THEN CSAA.TITLE ELSE (CS.TITLE + '-' + CSA.TITLE) END  AS PROPOSED_STATUS ,
		(SELECT SUM(CPP.DEFFERAL_DAYS + CL.DEFFERAL_DAYS) from ` + ReportTable + `CUSTOMER_POLICYCOMPLIANCELOG CL
		LEFT JOIN CUSTOMER_POLICYCOMPLIANCE CPP ON CPP.EXCEPTION_REFERENCEKEY = CL.EXCEPTION_REFERENCEKEY AND CPP.TRANSACTIONINPROGRESS_CODE=?1 AND CPP.EXCEPTION_REFERENCEKEY = CP.EXCEPTION_REFERENCEKEY
		WHERE CL.DEFERRALSTATUS=(Select STATUS_ID from CONFIG_STATUS where STATUS_KEY='A' ) AND CL.CUSTOMER_CODE = ?2 and CL.APPROVINGTRANSACTION_ID=?1) AS TOTAL_DAYS ,
		CP.EXCEPTION_COMMENTS as COMMENTS
		from CUSTOMERTRANSACTIONLOG CTL
		inner JOIN CUSTOMER_POLICYCOMPLIANCE CP ON CP.TRANSACTIONINPROGRESS_CODE = CTL.Transaction_id 
		LEFT JOIN CUSTOMER_POLICYCOMPLIANCELOG CL on CL.CUSTOMER_CODE = CTL.REFERENCE_ID and CL.EXCEPTION_REFERENCEKEY = CP.EXCEPTION_REFERENCEKEY
		LEFT JOIN CONFIG_POLICYANDEXCEPTIONSETUP CE ON CE.POLICYANDEXCEPTIONSETUP_ID = CL.POLICYANDEXCEPTIONSETUP_ID
		left JOIN ` + ReportTable + `CUSTOMER_POLICYCOMPLIANCELOG CLR ON CLR.TRANSACTIONINPROGRESS_CODE = CL.TRANSACTIONINPROGRESS_CODE and CLR.APPROVINGTRANSACTION_ID =?1  and CLR.EXCEPTION_REFERENCEKEY = CP.EXCEPTION_REFERENCEKEY
	
		inner JOIN CONFIG_ENTITYSTATUS CS ON CS.ENTITYSTATUS_ID=CP.EXCEPTIONENTITY_STATUS
		inner JOIN CONFIG_ENTITYSTATUS CSA ON CSA.ENTITYSTATUS_ID=CP.EXCEPTIONENTITY_STATUS_ACTION
		inner JOIN CONFIG_ENTITYSTATUS CSAA ON CSAA.ENTITYSTATUS_ID=CP.ENTITYSTATUS_ID
		inner JOIN ` + ReportTable + `CUSTOMER C ON C.TRANSACTIONINPROGRESS_CODE = ?1
		inner JOIN CONFIG_POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE PERS ON PERS.REGULATORYSEGMENT_ID=C.REGULATORYSEGMENT_ID and CE.POLICYANDEXCEPTIONSETUP_ID = PERS.POLICYANDEXCEPTIONSETUP_ID
		inner JOIN CONFIG_POLICYEXCEPTIONSETUPBUSINESSSEGMENTLINKAGES PEBS ON PEBS.BUSINESSHIERARCHY_ID=C.BUSINESSHIERARCHY_ID and CE.POLICYANDEXCEPTIONSETUP_ID = PEBS.POLICYANDEXCEPTIONSETUP_ID
	
		WHERE CL.DEFERRALSTATUS=(Select STATUS_ID from CONFIG_STATUS where STATUS_KEY='A' )
		AND CL.CUSTOMER_CODE = ?2 
		and CL.EXCEPTION_REFERENCEKEY IS NOT NULL
		and isnull(CE.POLICYANDEXCEPTIONSETUP_ID,0)= case when isnull(CE.POLICYANDEXCEPTIONSETUP_ID,0) !=0
		and ( PERS.POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE_ID is not null and PEBS.POLICYEXCEPTIONSETUPBUSINESS_ID is not null) then CE.POLICYANDEXCEPTIONSETUP_ID  when isnull(CE.POLICYANDEXCEPTIONSETUP_ID,0) =0 then 0 end
		and CP.STATUS = 1 and CSAA.ENTITYSTATUS_KEY='NC' and CTL.Transaction_id=?1
		
		UNION ALL
		
		Select  CPC.CUSTOMERPOLICYCOMPLIANCE_KEY AS REF,
		ISNULL(PES.TITLE,CPC.CUSTOMERPOLICYCOMPLIANCE_TITLE)  ,
		'' AS EXISTING_DEFFERAL_DAYS , '' AS EXISTING_DEFERRALCOUNT , '' AS EXISTING_DEFERRALEXPIRY ,
		ISNULL(CONVERT(varchar,CPC.DEFFERAL_DAYS),'')  AS PROPOSED_DEFFERAL_DAYS ,
		ISNULL(CONVERT(varchar,CPC.DEFERRALCOUNT),'')  AS PROPOSED_DEFERRALCOUNT ,
		FORMAT(CPC.DEFFERAL_UPTO,'MMMM dd, yyyy')  AS PROPOSED_DEFERRALEXPIRY ,
		ENS.TITLE + '-' + ENSA.TITLE AS PROPOSED_STATUS , CPC.DEFFERAL_DAYS  AS TOTAL_DAYS ,CPC.EXCEPTION_COMMENTS as COMMENTS
		from CUSTOMER_POLICYCOMPLIANCE CPC
		LEFT JOIN CONFIG_POLICYANDEXCEPTIONSETUP PES ON PES.POLICYANDEXCEPTIONSETUP_ID = CPC.POLICYANDEXCEPTIONSETUP_ID
		INNER JOIN CONFIG_ENTITYSTATUS ENS ON ENS.ENTITYSTATUS_ID=CPC.EXCEPTIONENTITY_STATUS AND ENS.ENTITYSTATUS_KEY='Defered'
		INNER JOIN CONFIG_ENTITYSTATUS ENTS ON ENTS.ENTITYSTATUS_ID = CPC.ENTITYSTATUS_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ENSA ON ENSA.ENTITYSTATUS_ID=CPC.EXCEPTIONENTITY_STATUS_ACTION
		LEFT JOIN ` + ReportTable + `CUSTOMER C ON C.TRANSACTIONINPROGRESS_CODE = CPC.TRANSACTIONINPROGRESS_CODE
		LEFT JOIN CONFIG_POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE PERS ON PERS.REGULATORYSEGMENT_ID=C.REGULATORYSEGMENT_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PERS.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_POLICYEXCEPTIONSETUPBUSINESSSEGMENTLINKAGES PEBS ON PEBS.BUSINESSHIERARCHY_ID=C.BUSINESSHIERARCHY_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PEBS.POLICYANDEXCEPTIONSETUP_ID
		WHERE CPC.TRANSACTIONINPROGRESS_CODE=?1 AND CPC.TRANSACTION_CODE=?1 AND CPC.EXCEPTION_REFERENCEKEY IS NULL
		and isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0)= case when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) !=0  and ( PERS.POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE_ID is not null and PEBS.POLICYEXCEPTIONSETUPBUSINESS_ID is not null) then PES.POLICYANDEXCEPTIONSETUP_ID  when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) =0 then 0 end
		and CPC.STATUS = 1 and ENTS.ENTITYSTATUS_KEY='NC'`

	txRows, err := Db.Query(query, transId, referenceId, transId, transId, referenceId, transId, transId, transId)
	if err != nil {
		log.Printf("Error executing GetDeferralDetails query: %v", err)
		return nil
	}
	defer txRows.Close()

	var results []CMS.DeferalWaiversOpenExceptionDetails
	for txRows.Next() {
		var item CMS.DeferalWaiversOpenExceptionDetails
		err = txRows.Scan(&item.Ref, &item.TITLE, &item.EXISTING_DEFFERAL_DAYS, &item.EXISTING_DEFERRALCOUNT, &item.EXISTING_DEFERRALEXPIRY,
			&item.PROPOSED_DEFFERAL_DAYS, &item.PROPOSED_DEFERRALCOUNT, &item.PROPOSED_DEFERRALEXPIRY, &item.PROPOSED_STATUS,
			&item.TOTAL_DAYS, &item.Comments)
		if err != nil {
			log.Printf("Error scanning GetDeferralDetails row: %v", err)
			continue
		}
		results = append(results, item)
	}

	if err = txRows.Err(); err != nil {
		log.Printf("Row iteration error in GetDeferralDetails: %v", err)
	}

	return results
}

func GetWaiversDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.DeferalWaiversOpenExceptionDetails {

	query := `SELECT ISNULL(CPC.CUSTOMERPOLICYCOMPLIANCE_KEY,'-') AS id,ISNULL(PES.TITLE,CPC.CUSTOMERPOLICYCOMPLIANCE_TITLE) AS description,
		ISNULL(ENSA.TITLE,'-') AS status,ISNULL(CPC.EXCEPTION_COMMENTS,'-') AS comments
		FROM CUSTOMER_POLICYCOMPLIANCE CPC
		LEFT JOIN CONFIG_POLICYANDEXCEPTIONSETUP PES ON PES.POLICYANDEXCEPTIONSETUP_ID = CPC.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ENTS ON ENTS.ENTITYSTATUS_ID = CPC.ENTITYSTATUS_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ENS ON ENS.ENTITYSTATUS_ID = CPC.EXCEPTIONENTITY_STATUS
		LEFT JOIN CONFIG_ENTITYSTATUS ENSA ON ENSA.ENTITYSTATUS_ID = CPC.EXCEPTIONENTITY_STATUS_ACTION
		LEFT JOIN ` + ReportTable + `CUSTOMER C ON C.TRANSACTIONINPROGRESS_CODE = CPC.TRANSACTIONINPROGRESS_CODE
		LEFT JOIN CONFIG_POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE PERS ON PERS.REGULATORYSEGMENT_ID=C.REGULATORYSEGMENT_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PERS.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_POLICYEXCEPTIONSETUPBUSINESSSEGMENTLINKAGES PEBS ON PEBS.BUSINESSHIERARCHY_ID=C.BUSINESSHIERARCHY_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PEBS.POLICYANDEXCEPTIONSETUP_ID
		WHERE CPC.TRANSACTIONINPROGRESS_CODE =?1
		AND ENS.ENTITYSTATUS_KEY = 'Waived'
		and isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0)= case when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) !=0  and ( PERS.POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE_ID is not null and PEBS.POLICYEXCEPTIONSETUPBUSINESS_ID is not null) then PES.POLICYANDEXCEPTIONSETUP_ID  when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) =0 then 0 end
		and CPC.STATUS = 1  and ENTS.ENTITYSTATUS_KEY='NC'`

	txRows, err := Db.Query(query, transId)
	if err != nil {
		log.Printf("Error executing GetWaiversDetails query: %v", err)
		return nil
	}
	defer txRows.Close()

	var results []CMS.DeferalWaiversOpenExceptionDetails
	for txRows.Next() {
		var item CMS.DeferalWaiversOpenExceptionDetails
		err = txRows.Scan(&item.Ref, &item.Description, &item.Status, &item.Comments)
		if err != nil {
			log.Printf("Error scanning GetWaiversDetails row: %v", err)
			continue
		}
		results = append(results, item)
	}

	if err = txRows.Err(); err != nil {
		log.Printf("Row iteration error in GetWaiversDetails: %v", err)
	}

	return results
}

func GetOpenExceptionDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.DeferalWaiversOpenExceptionDetails {

	query := `SELECT ISNULL(CPC.CUSTOMERPOLICYCOMPLIANCE_KEY,'-') AS id,ISNULL(PES.TITLE,CPC.CUSTOMERPOLICYCOMPLIANCE_TITLE) AS description,
		ISNULL(ENSA.TITLE,'-') AS status,ISNULL(CPC.EXCEPTION_COMMENTS,'-') AS comments
		FROM CUSTOMER_POLICYCOMPLIANCE CPC
		LEFT JOIN CONFIG_POLICYANDEXCEPTIONSETUP PES ON PES.POLICYANDEXCEPTIONSETUP_ID = CPC.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ENTS ON ENTS.ENTITYSTATUS_ID = CPC.ENTITYSTATUS_ID
		LEFT JOIN CONFIG_ENTITYSTATUS ENS ON ENS.ENTITYSTATUS_ID = CPC.EXCEPTIONENTITY_STATUS
		LEFT JOIN CONFIG_ENTITYSTATUS ENSA ON ENSA.ENTITYSTATUS_ID = CPC.EXCEPTIONENTITY_STATUS_ACTION
		LEFT JOIN ` + ReportTable + `CUSTOMER C ON C.TRANSACTIONINPROGRESS_CODE = CPC.TRANSACTIONINPROGRESS_CODE
		LEFT JOIN CONFIG_POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE PERS ON PERS.REGULATORYSEGMENT_ID=C.REGULATORYSEGMENT_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PERS.POLICYANDEXCEPTIONSETUP_ID
		LEFT JOIN CONFIG_POLICYEXCEPTIONSETUPBUSINESSSEGMENTLINKAGES PEBS ON PEBS.BUSINESSHIERARCHY_ID=C.BUSINESSHIERARCHY_ID and PES.POLICYANDEXCEPTIONSETUP_ID = PEBS.POLICYANDEXCEPTIONSETUP_ID
		WHERE CPC.TRANSACTIONINPROGRESS_CODE =?1
		AND ENS.ENTITYSTATUS_KEY = 'open'
		and isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0)= case when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) !=0  and ( PERS.POLICYEXCEPTIONREGULATORYSEGMENTLINKAGE_ID is not null and PEBS.POLICYEXCEPTIONSETUPBUSINESS_ID is not null) then PES.POLICYANDEXCEPTIONSETUP_ID  when isnull(PES.POLICYANDEXCEPTIONSETUP_ID,0) =0 then 0 end
		and CPC.STATUS = 1 and ENTS.ENTITYSTATUS_KEY='NC'`

	txRows, err := Db.Query(query, transId)
	if err != nil {
		log.Printf("Error executing GetOpenExceptionDetails query: %v", err)
		return nil
	}
	defer txRows.Close()

	var results []CMS.DeferalWaiversOpenExceptionDetails
	for txRows.Next() {
		var item CMS.DeferalWaiversOpenExceptionDetails
		err = txRows.Scan(&item.Ref, &item.Description, &item.Status, &item.Comments)
		if err != nil {
			log.Printf("Error scanning GetOpenExceptionDetails row: %v", err)
			continue
		}
		results = append(results, item)
	}

	if err = txRows.Err(); err != nil {
		log.Printf("Row iteration error in GetOpenExceptionDetails: %v", err)
	}

	return results
}

func GetFacilityDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.FacilityDetails {

	query := `SELECT ISNULL(CF.SEQUENCE,'-') AS line,
CASE WHEN CF.OTHERGROUP_CUSTOMERCHILD = 1 THEN (SELECT CONCAT(f.TITLE ,' <em>(Sublimit of ' , CFSA.SEQUENCE, ' - ', fapp.TITLE  , ' of ' , capp.TITLE , ')</em>')
FROM CUSTOMERTRANSACTIONLOG CTLP
LEFT JOIN CUSTOMER_FACILITYAPPROVED CFP on CFP.CUSTOMER_ID = CTLP.REFERENCE_ID 
left JOIN CUSTOMER_FACILITYAPPROVED CFSA ON CFSA.CUSTOMERFACILITY_ID = CFP.PARENTID 
left join CONFIG_FACILITY fapp on fapp.FACILITY_ID = CFSA.FACILITY_ID
left join CUSTOMERAPPROVED capp on capp.CUSTOMER_ID = CFSA.CUSTOMER_ID 
WHERE CTLP.TRANSACTION_ID=?1 AND CFP.CUSTOMERFACILITY_ID = CF.CUSTOMERFACILITY_ID)
ELSE f.TITLE END AS facilitydetails,
ISNULL(BH.TITLE,'-') AS controlunit,
ISNULL(AA.TITLE,'-') AS request,
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(CFS.TITLE,'-') ELSE '' end as secured,
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(CF.AVAILABILITYPERIOD,'-') else '' end as availabilityperiod,
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN CASE WHEN FT.FACILITYTYPE_KEY  = 'F' THEN
(CASE WHEN ISNULL(BM.TITLE,'') ='' OR ISNULL(CF.PRICING + '%','') ='' THEN '-'
ELSE ISNULL(ISNULL(BM.TITLE,'') + CASE WHEN CF.PRICING LIKE '-%' THEN ' ' ELSE ' + ' END + ISNULL(CF.PRICING + '%',''),'-')  END) ELSE ISNULL(CF.COMMISSIONASPERLOC,'') END else '' end as pricing,
ISNULL(FORMAT(CF.FACILITYEXPIRY, 'MMMM dd, yyyy'),'-') as expirydate,
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(Convert(varchar,CONVERT(BIGINT,CF.TENOR_MONTHS)) + ' Months','-') else '' end as maxtenor,
ISNULL(CFA.PARENTID ,'-') AS exctparentid,
ISNULL(CF.FACILITYLIMIT,'') as proposedlimit,ISNULL(CFA.FACILITYLIMIT,'') as existinglimit,ISNULL(CF.PARENTID,'-') AS parentid,
ISNULL(RS.TITLE,'-') as proposedfrr, ISNULL(RSA.TITLE,'-') as existingfrr,

proposedsecurities=ISNULL(replace(replace(STUFF( (
SELECT * FROM (select CONCAT(ISNULL(sec.TITLE,''),' of ',
ISNULL(FORMAT(CONVERT(bigint,CSCH.CHARGEAMOUNT),'#,###,##0'),'0'),' with ',
ISNULL(CONVERT(float,cs.MARGINPERCENTAGE),''),'% Margin - (', CFTS.TITLE,')') as span
from CUSTOMERTRANSACTIONLOG CTLFL
LEFT JOIN CUSTOMER_FACILITYCOLLATERALLINKAGESAPPROVED cfl on cfl.CUSTOMER_CODE = CTLFL.REFERENCE_ID
LEFT JOIN FaysalBank_DataLogs_CMS.dbo.REPORT_CUSTOMER_SECURITY cs  on cs.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE 
LEFT JOIN CONFIG_SECURITY sec on sec.SECURITY_ID = cs.SECURITY_CODE
LEFT JOIN CONFIG_PRIORITYTYPE cfts on cfts.PRIORITYTYPE_ID  = cfl.COLLATERAL_PRIORITY
LEFT JOIN CUSTOMER_SECURITYCHARGEDETAILLINKAGEAPPROVED CSCH ON CSCH.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE
AND CSCH.BANKSETUP = (SELECT BANKSETUP FROM CONFIG_BANKSETUP WHERE OTHERBANK = 0) AND CSCH.STATUS_ID=1

where CTLFL.CUSTOMER_ID=?2 AND cfl.CUSTOMERFACILITY_CODE= CF.CUSTOMERFACILITY_ID AND cfl.LINK =1 and CS.IS_CHARGEABLECOLLATERAL = 1

UNION ALL

select CONCAT(ISNULL(sec.TITLE,''),' of ',
ISNULL(FORMAT(CONVERT(bigint,CSCH.CHARGEAMOUNT),'#,###,##0'),'0'),' with ',
ISNULL(CONVERT(float,cs.MARGINPERCENTAGE),''),'% Margin - (', CFTS.TITLE,')') as span
from CUSTOMERTRANSACTIONLOG CTLFL
LEFT JOIN CUSTOMER_FACILITYCOLLATERALLINKAGESAPPROVED cfl on cfl.CUSTOMER_CODE = CTLFL.REFERENCE_ID 
LEFT JOIN CUSTOMER_SECURITYAPPROVED cs  on cs.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE
LEFT JOIN CONFIG_PRIORITYTYPE cfts on cfts.PRIORITYTYPE_ID  = cfl.COLLATERAL_PRIORITY
LEFT JOIN CONFIG_SECURITY sec on sec.SECURITY_ID  = cs.SECURITY_CODE
LEFT JOIN CUSTOMER_SECURITYCHARGEDETAILLINKAGEAPPROVED CSCH ON CSCH.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE 
AND CSCH.BANKSETUP = (SELECT BANKSETUP FROM CONFIG_BANKSETUP WHERE OTHERBANK = 0) AND CSCH.STATUS_ID=1
where CTLFL.TRANSACTION_ID=?1 AND cfl.CUSTOMERFACILITY_CODE= CF.CUSTOMERFACILITY_ID AND cfl.LINK =1 and CS.IS_CHARGEABLECOLLATERAL = 1 and cfl.IS_GROUP =1  and cs.status =1) t
FOR XML PATH ('li')),1,0,'' ),'&lt;','<'),'&gt;','>'),'-'),

CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN CASE WHEN FT.FACILITYTYPE_KEY = 'F' THEN ISNULL(CONCAT('<p><strong>Principal Repayment Terms:</strong> ' , ISNULL(CF.PRINCIPALREPAYMENTTERMS,'') , ' </p><p><strong>Profit Payment Terms:</strong> ' ,
ISNULL(CF.PRICINGREPAYMENTTERMS,'') ,' </p><p><strong>Repayment Frequency:</strong> ',ISNULL(FSQ.TITLE,''),'</p>'),'-')
ELSE ISNULL(CF.COMMISSIONREPAYMENTTERM,'-')END  else '' end as repaymentTerms,
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(CF.REASONFORCHANGE,'-') else '' end as purpose,CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(CF.PROPOSEDFACILITYSPECIFICCONDITION,'-') else '' end as condition,

existingsecurities= ISNULL(replace(replace(STUFF( (
SELECT * FROM (select CONCAT(ISNULL(sec.TITLE,''),' of ',
ISNULL(FORMAT(CONVERT(bigint,CSCH.CHARGEAMOUNT),'#,###,##0'),'0'),' with ',
ISNULL(CONVERT(float,cs.MARGINPERCENTAGE),''),'% Margin - (', CFTS.TITLE,')') as span
from CUSTOMERTRANSACTIONLOG CTLFL
LEFT JOIN CUSTOMER_FACILITYCOLLATERALLINKAGESAPPROVED cfl on cfl.CUSTOMER_CODE = CTLFL.REFERENCE_ID
LEFT JOIN CUSTOMER_SECURITYAPPROVED cs  on cs.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE 
LEFT JOIN CONFIG_SECURITY sec on sec.SECURITY_ID = cs.SECURITY_CODE
LEFT JOIN CONFIG_PRIORITYTYPE cfts on cfts.PRIORITYTYPE_ID  = cfl.COLLATERAL_PRIORITY
LEFT JOIN CUSTOMER_SECURITYCHARGEDETAILLINKAGEAPPROVED CSCH ON CSCH.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE
AND CSCH.BANKSETUP = (SELECT BANKSETUP FROM CONFIG_BANKSETUP WHERE OTHERBANK = 0) AND CSCH.STATUS_ID=1

where CTLFL.CUSTOMER_ID=?2 AND cfl.CUSTOMERFACILITY_CODE= CF.CUSTOMERFACILITY_ID AND cfl.LINK =1 and CS.IS_CHARGEABLECOLLATERAL = 1
UNION ALL

select CONCAT(ISNULL(sec.TITLE,''),' of ',
ISNULL(FORMAT(CONVERT(bigint,CSCH.CHARGEAMOUNT),'#,###,##0'),'0'),' with ',
ISNULL(CONVERT(float,cs.MARGINPERCENTAGE),''),'% Margin - (', CFTS.TITLE,')') as span
from CUSTOMERTRANSACTIONLOG CTLFL
LEFT JOIN CUSTOMER_FACILITYCOLLATERALLINKAGESAPPROVED cfl on cfl.CUSTOMER_CODE = CTLFL.REFERENCE_ID
LEFT JOIN CUSTOMER_SECURITYAPPROVED cs  on cs.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE 
LEFT JOIN CONFIG_PRIORITYTYPE cfts on cfts.PRIORITYTYPE_ID  = cfl.COLLATERAL_PRIORITY
LEFT JOIN CONFIG_SECURITY sec on sec.SECURITY_ID  = cs.SECURITY_CODE
LEFT JOIN CUSTOMER_SECURITYCHARGEDETAILLINKAGEAPPROVED CSCH ON CSCH.CUSTOMERSECURITY_ID = cfl.CUSTOMERSECURITY_CODE
AND CSCH.BANKSETUP = (SELECT BANKSETUP FROM CONFIG_BANKSETUP WHERE OTHERBANK = 0) AND CSCH.STATUS_ID=1
where CTLFL.CUSTOMER_ID=?2 AND cfl.CUSTOMERFACILITY_CODE= CF.CUSTOMERFACILITY_ID AND cfl.LINK =1 and CS.IS_CHARGEABLECOLLATERAL = 1and cfl.IS_GROUP =1  and cs.status =1) t
FOR XML PATH ('li')),1,0,'' ),'&lt;','<'),'&gt;','>'),'-'),
CASE WHEN AA.FACILITYSTATUS_KEY != 'CANCELLATION' THEN ISNULL(CF.REPAYMENTHISTORY,'-') else '' end as repaymenthistory

FROM CUSTOMERTRANSACTIONLOG CTL
LEFT JOIN CUSTOMER_FACILITYAPPROVED CF on CF.CUSTOMER_ID = ctl.REFERENCE_ID 
LEFT JOIN CONFIG_FACILITY F ON F.FACILITY_ID =  CF.FACILITY_ID
LEFT JOIN CONFIG_RATINGSETUP RS ON RS.RATINGSETUP_ID = CF.FRR
LEFT join CONFIG_STATUS as st on st.STATUS_ID=  CF.STATUS
LEFT JOIN CONFIG_FACILITYSTATUS FS ON FS.FACILITYSTATUS_ID =  CF.FACILITYSTATUS_ID
LEFT JOIN CONFIG_BUSINESSHIERARCHY BH ON BH.BUSINESSHIERARCHY_ID = CTL.BUSINESSHIERARCHY_ID
LEFT JOIN CONFIG_FACILITYCLASSIFICATION CS ON CS.FACILITYCLASSIFICATION_ID = F.CLASSIFICATION_ID
LEFT JOIN CONFIG_FACILITYSECURETYPE CFS ON CFS.FACILITYSECURETYPE_ID =  CF.FACILITYSECURETYPE_ID
LEFT JOIN CONFIG_TENORTYPE TE on TE.TENORTYPE_ID =  CF.TENORTYPE_ID
LEFT JOIN CONFIG_BENCHMARKSETUP BM on BM.BENCHMARKSETUP_ID =  CF.BENCHMARKSETUP
LEFT JOIN CONFIG_FREQUENCYSETUP FSQ on FSQ.FREQUENCYSETUP_ID =  CF.REPAYMENTFREQUENCY
LEFT JOIN CONFIG_FACILITYSTATUS AA ON AA.FACILITYSTATUS_ID =  CF.FACILITYSTATUS_ID
LEFT JOIN CONFIG_FREQUENCYSETUP RPF on RPF.FREQUENCYSETUP_ID = CF.REPRICINGFREQUENCY
LEFT JOIN CONFIG_FACILITYTYPE FT on FT.FACILITYTYPE_ID = F.FACILITYTYPE_ID
LEFT JOIN CUSTOMER_FACILITYAPPROVED CFA ON CFA.REFERENCE_ID =  CF.REFERENCE_ID
LEFT JOIN CONFIG_RATINGSETUP RSA ON RSA.RATINGSETUP_ID = CFA.FRR
WHERE CTL.TRANSACTION_ID=?1 and st.STATUS_KEY='A' 
ORDER BY CAST('/'+REPLACE( CF.sequence,'.','/')+'/' as hierarchyID)`

	rows, err := Db.Query(query, transId, referenceId)
	if err != nil {
		log.Printf("Error executing GetFacilityDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var facilities []CMS.FacilityDetails
	for rows.Next() {
		var f CMS.FacilityDetails
		err := rows.Scan(
			&f.Line,
			&f.FacilityDetails,
			&f.ControlUnit,
			&f.Request,
			&f.Secured,
			&f.AvailabilityPeriod,
			&f.Pricing,
			&f.FacilityExpiry,
			&f.MaxTenor,
			&f.ExistingParentId,
			&f.ProposedLimit,
			&f.ExistingLimit,
			&f.ParentId,
			&f.ProposedFRR,
			&f.ExistingFRR,
			&f.ProposedSecurities,
			&f.RepaymentTerms,
			&f.Purpose,
			&f.Condition,
			&f.ExistingSecurities,
			&f.RepaymentHistory)
		if err != nil {
			log.Printf("Error scanning GetFacilityDetails row: %v", err)
			continue
		}
		facilities = append(facilities, f)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetFacilityDetails: %v", err)
	}

	return facilities
}

func GetApproversCommentsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.ApproversCommentsDetails {

	query := ` select t.Designation,t.Name,t.Comments,t.DATE,t.actionTaken from (Select  w.WORKFLOWLOG_ID,ws.title as Designation, us.Name as Name, w.COMMENT as Comments, 
		FORMAT(w.DATE,'MMMM dd, yyyy hh:mm:ss tt') as DATE,	case when w.casereinitiated=1 then 'Reinitiated' else	wsp.title end as actionTaken
		from WF_WORKFLOWLOG w
		inner join WF_WORKFLOWSTATE ws on ws.WORKFLOWSTATE_ID = w.WORKFLOWSTATE_ID
		inner join WF_WORKFLOWSTEP wsp on wsp.WORKFLOWSTEP_ID = w.WORKFLOWSTEP_ID
		left join CONFIG_USERS us on us.USER_ID = w.EMPLOYEE_ID
		where w.RECORD_ID = ?1
		union 
		select w.WORKFLOWLOG_ID,ws.title as Designation, us.Name as Name, w.COMMENT as Comments, FORMAT(w.DATE,'MMMM dd, yyyy hh:mm:ss tt') as DATE,		
		case when w.casereinitiated=1 then 'Reinitiated' else	ws.title end as actionTaken
		from WF_WORKFLOWLOG w 
		inner join WF_WORKFLOWSTATE ws on ws.WORKFLOWSTATE_ID = w.WORKFLOWSTATE_ID
		left join CONFIG_USERS us on us.USER_ID = w.EMPLOYEE_ID
		where w.RECORD_ID = ?1 and w.casereinitiated=1) t  order by t.workflowlog_id desc`
	rows, err := Db.Query(query, transId)
	if err != nil {
		log.Printf("Error executing GetApproversCommentsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var approversComments []CMS.ApproversCommentsDetails
	for rows.Next() {
		var item CMS.ApproversCommentsDetails
		err = rows.Scan(&item.Designation, &item.Name, &item.Comments, &item.Date, &item.ActionTaken)
		if err != nil {
			log.Printf("Error scanning GetApproversCommentsDetails row: %v", err)
			continue
		}
		approversComments = append(approversComments, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetApproversCommentsDetails: %v", err)
	}

	return approversComments
}

func GetTotalAssetsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) CMS.TotalAssetsDetails {

	//var rows *sql.Rows
	//var totalAssetsList []CMS.TotalAssetsDetails
	//for rows.Next() {
	var totalAssets CMS.TotalAssetsDetails
	currentAssets := GetCurrentAssetsDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.CurrentAssets = currentAssets

	fixedAssets := GetFixedAssetsDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.FixedAssets = fixedAssets

	mortgage := GetMortgageDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.Mortgage = mortgage

	stocks := GetStocksHeldUnderPledgeDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.StocksHeldUnderPledge = stocks

	pledge := GetPledgeOfSharesDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.PledgeOfShares = pledge

	gop := GetGopDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.Gop = gop

	cashCollateral := GetCashCollateralDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.CashCollateral = cashCollateral

	collateralSummary := GetCollateralSummaryDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.CollateralSummary = collateralSummary

	totalcurrentAssets := GetTotalCurrentAssetsDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalCurrentAssets = totalcurrentAssets

	totalfixedAssets := GetTotalFixedAssetsDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalFixedAssets = totalfixedAssets

	totalmortgage := GetTotalMortgageDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalMortgage = totalmortgage

	totalstocks := GetTotalStocksHeldUnderPledgeDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalStocksHeldUnderPledge = totalstocks

	totalpledge := GetTotalPledgeOfSharesDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalPledgeOfShares = totalpledge

	totalgop := GetTotalGopDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalGop = totalgop

	totalcashCollateral := GetTotalCashCollateralDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalCashCollateral = totalcashCollateral

	totalcollateralSummary := GetTotalCollateralSummaryDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	totalAssets.TotalCollateralSummary = totalcollateralSummary

	//	totalAssetsList = append(totalAssetsList, totalAssets)
	//}
	return totalAssets
}

func GetCurrentAssetsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 1 as id, 'Current Assets' as Item, ISNULL(FORMAT(5000,'###,###,##0'),'-') as amount, 'secured' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetCurrentAssetsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var currentAssets []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetCurrentAssetsDetails row: %v", err)
			continue
		}
		currentAssets = append(currentAssets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetCurrentAssetsDetails: %v", err)
	}

	return currentAssets
}

func GetFixedAssetsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 2 as id, 'Fixed Assets' as Item, ISNULL(FORMAT(6000,'###,###,##0'),'-') as amount, 'secured' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetFixedAssetsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var fixedAssets []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetFixedAssetsDetails row: %v", err)
			continue
		}
		fixedAssets = append(fixedAssets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetFixedAssetsDetails: %v", err)
	}

	return fixedAssets
}

func GetMortgageDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 3 as id, 'Mortgage' as Item, ISNULL(FORMAT(7000,'###,###,##0'),'-') as amount, 'secured' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetMortgageDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var mortgage []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetMortgageDetails row: %v", err)
			continue
		}
		mortgage = append(mortgage, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetMortgageDetails: %v", err)
	}

	return mortgage
}

func GetStocksHeldUnderPledgeDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 4 as id, 'Stocks Held Under Pledge' as Item, ISNULL(FORMAT(8000,'###,###,##0'),'-') as amount, 'secured' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetStocksHeldUnderPledgeDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var stocks []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetStocksHeldUnderPledgeDetails row: %v", err)
			continue
		}
		stocks = append(stocks, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetStocksHeldUnderPledgeDetails: %v", err)
	}

	return stocks
}

func GetCashCollateralDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 5 as id, 'Cash Collateral' as Item, ISNULL(FORMAT(9000,'###,###,##0'),'-') as amount, 'N/A' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetCashCollateralDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var cashCollateral []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetCashCollateralDetails row: %v", err)
			continue
		}
		cashCollateral = append(cashCollateral, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetCashCollateralDetails: %v", err)
	}

	return cashCollateral
}

func GetGopDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 6 as id, 'Gop' as Item, ISNULL(FORMAT(10000,'###,###,##0'),'-') as amount, 'N/A' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetGopDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var gop []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetGopDetails row: %v", err)
			continue
		}
		gop = append(gop, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetGopDetails: %v", err)
	}

	return gop
}

func GetPledgeOfSharesDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChild {

	query := `SELECT 7 as id, 'Pledge of Shares' as Item, ISNULL(FORMAT(11000,'###,###,##0'),'-') as amount, 'N/A' as status
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetPledgeOfSharesDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var pledge []CMS.TotalAssetsChild
	for rows.Next() {
		var item CMS.TotalAssetsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Amount, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetPledgeOfSharesDetails row: %v", err)
			continue
		}
		pledge = append(pledge, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetPledgeOfSharesDetails: %v", err)
	}

	return pledge
}

func GetCollateralSummaryDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.CollateralSummary {

	query := `SELECT ISNULL(FORMAT(5000,'###,###,##0'),'-') as currentassetamount, 'secured' as currentassetstatus, 
	ISNULL(FORMAT(6000,'###,###,##0'),'-') as fixedassetamount, 'secured' as fixedassetstatus,
	ISNULL(FORMAT(7000,'###,###,##0'),'-') as mortgageamount, 'secured' as mortgagestatus,
	ISNULL(FORMAT(8000,'###,###,##0'),'-') as stockamount, 'N/A' as stockstatus,
	ISNULL(FORMAT(9000,'###,###,##0'),'-') as cashcollateralamount, 'N/A' as cashcollateralstatus,
	ISNULL(FORMAT(10000,'###,###,##0'),'-') as gopamount, 'N/A' as gopstatus,
	ISNULL(FORMAT(11000,'###,###,##0'),'-') as pledgeamount, 'N/A' as pledgestatus
	--FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetCollateralSummaryDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var collateralSummary []CMS.CollateralSummary
	for rows.Next() {
		var item CMS.CollateralSummary
		err = rows.Scan(&item.CurrentAssetAmount, &item.CurrentAssetStatus,
			&item.FixedAssetAmount, &item.FixedAssetStatus,
			&item.MortgageAmount, &item.MortgageStatus,
			&item.StockAmount, &item.StockStatus,
			&item.CashCollateralAmount, &item.CashCollateralStatus,
			&item.GopAmount, &item.GopStatus,
			&item.PledgeAmount, &item.PledgeStatus)
		if err != nil {
			log.Printf("Error scanning GetCollateralSummaryDetails row: %v", err)
			continue
		}
		collateralSummary = append(collateralSummary, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetCollateralSummaryDetails: %v", err)
	}

	return collateralSummary
}

func GetTotalCurrentAssetsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(5000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalCurrentAssetsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var currentAssets []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalCurrentAssetsDetails row: %v", err)
			continue
		}
		currentAssets = append(currentAssets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalCurrentAssetsDetails: %v", err)
	}

	return currentAssets
}

func GetTotalFixedAssetsDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(6000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalFixedAssetsDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var fixedAssets []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalFixedAssetsDetails row: %v", err)
			continue
		}
		fixedAssets = append(fixedAssets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalFixedAssetsDetails: %v", err)
	}

	return fixedAssets
}

func GetTotalMortgageDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(7000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalMortgageDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var mortgage []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalMortgageDetails row: %v", err)
			continue
		}
		mortgage = append(mortgage, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalMortgageDetails: %v", err)
	}

	return mortgage
}

func GetTotalStocksHeldUnderPledgeDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(8000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalStocksHeldUnderPledgeDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var stocks []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalStocksHeldUnderPledgeDetails row: %v", err)
			continue
		}
		stocks = append(stocks, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalStocksHeldUnderPledgeDetails: %v", err)
	}

	return stocks
}

func GetTotalCashCollateralDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(9000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalCashCollateralDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var cashcollateral []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalCashCollateralDetails row: %v", err)
			continue
		}
		cashcollateral = append(cashcollateral, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalCashCollateralDetails: %v", err)
	}

	return cashcollateral
}

func GetTotalGopDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(10000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalGopDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var gop []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalGopDetails row: %v", err)
			continue
		}
		gop = append(gop, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalGopDetails: %v", err)
	}

	return gop
}

func GetTotalPledgeOfSharesDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalAssetsChildTotal {

	query := `SELECT ISNULL(FORMAT(SUM(11000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalPledgeOfSharesDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var pledge []CMS.TotalAssetsChildTotal
	for rows.Next() {
		var item CMS.TotalAssetsChildTotal
		err = rows.Scan(&item.TotalAmount)
		if err != nil {
			log.Printf("Error scanning GetTotalPledgeOfSharesDetails row: %v", err)
			continue
		}
		pledge = append(pledge, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalPledgeOfSharesDetails: %v", err)
	}

	return pledge
}

func GetTotalCollateralSummaryDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.TotalCollateralSummary {

	query := `SELECT ISNULL(FORMAT(SUM(56000),'###,###,##0'),'-') as amount
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTotalCollateralSummaryDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var collateralvalue []CMS.TotalCollateralSummary
	for rows.Next() {
		var item CMS.TotalCollateralSummary
		err = rows.Scan(&item.TotalCollateralValue)
		if err != nil {
			log.Printf("Error scanning GetTotalCollateralSummaryDetails row: %v", err)
			continue
		}
		collateralvalue = append(collateralvalue, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTotalCollateralSummaryDetails: %v", err)
	}

	return collateralvalue
}

func GetMonitoringItemDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) CMS.MonitoringItems {

	//var rows *sql.Rows
	//var totalAssetsList []CMS.TotalAssetsDetails
	//for rows.Next() {
	var monitoringitems CMS.MonitoringItems
	takaful := GetTakafulMonitoringDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.Takaful = takaful

	fixedAssets := GetFixedAssetsExpiresDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.FixedAssets = fixedAssets

	stocks := GetInspectionOfStocksDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.Stocks = stocks

	currentassets := GetCurrentAssetsReportDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.CurrentAssets = currentassets

	pledge := GetPledgeStocksDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.Pledge = pledge

	monitoringsummary := GetMonitoringSummaryDetails(Db, DbApproved, transId, referenceId, alertActivityTypeKey, fromApproved, ReportTable)
	monitoringitems.MonitoringSummary = monitoringsummary

	//	totalAssetsList = append(totalAssetsList, totalAssets)
	//}
	return monitoringitems
}

func GetTakafulMonitoringDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringItemsChild {

	query := `SELECT 1 as id, 'TAKAFUL MONITORING' as Item, 'No Breach' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetTakafulMonitoringDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var takaful []CMS.MonitoringItemsChild
	for rows.Next() {
		var item CMS.MonitoringItemsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetTakafulMonitoringDetails row: %v", err)
			continue
		}
		takaful = append(takaful, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetTakafulMonitoringDetails: %v", err)
	}

	return takaful

}

func GetFixedAssetsExpiresDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringItemsChild {

	query := `SELECT 3 as id, 'FIXED ASSETS EXPIRIES' as Item, 'Breach' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetFixedAssetsExpiresDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var fixedassets []CMS.MonitoringItemsChild
	for rows.Next() {
		var item CMS.MonitoringItemsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetFixedAssetsExpiresDetails row: %v", err)
			continue
		}
		fixedassets = append(fixedassets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetFixedAssetsExpiresDetails: %v", err)
	}

	return fixedassets
}

func GetInspectionOfStocksDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringItemsChild {

	query := `SELECT 2 as id, 'INSPECTION OF STOCKS' as Item, 'No Breach' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetInspectionOfStocksDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var stocks []CMS.MonitoringItemsChild
	for rows.Next() {
		var item CMS.MonitoringItemsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetInspectionOfStocksDetails row: %v", err)
			continue
		}
		stocks = append(stocks, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetInspectionOfStocksDetails: %v", err)
	}

	return stocks
}

func GetCurrentAssetsReportDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringItemsChild {

	query := `SELECT 4 as id, 'CURRENT ASSETS REPORT' as Item, 'Breach' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetCurrentAssetsReportDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var currentassets []CMS.MonitoringItemsChild
	for rows.Next() {
		var item CMS.MonitoringItemsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetCurrentAssetsReportDetails row: %v", err)
			continue
		}
		currentassets = append(currentassets, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetCurrentAssetsReportDetails: %v", err)
	}

	return currentassets
}

func GetPledgeStocksDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringItemsChild {

	query := `SELECT 5 as id, 'PLEDGE STOCKS' as Item, 'No Breach' as status
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetPledgeStocksDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var pledge []CMS.MonitoringItemsChild
	for rows.Next() {
		var item CMS.MonitoringItemsChild
		err = rows.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			log.Printf("Error scanning GetPledgeStocksDetails row: %v", err)
			continue
		}
		pledge = append(pledge, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetPledgeStocksDetails: %v", err)
	}

	return pledge
}

func GetMonitoringSummaryDetails(Db *sql.DB, DbApproved *sql.DB, transId int, referenceId int, alertActivityTypeKey string, fromApproved bool, ReportTable string) []CMS.MonitoringSummary {

	query := `SELECT 'No Breach' as takaful,'No Breach' as stocks,'No Breach' as fixedassets,'No Breach' as currentassets,'No Breach' as pledge
	---FROM ` + ReportTable + `CUSTOMER_SECURITY CS
	--WHERE CS.TRANSACTIONINPROGRESS_CODE =?1`

	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error executing GetMonitoringSummaryDetails query: %v", err)
		return nil
	}
	defer rows.Close()

	var monitoringsummary []CMS.MonitoringSummary
	for rows.Next() {
		var item CMS.MonitoringSummary
		err = rows.Scan(&item.Takaful, &item.Stocks, &item.FixedAssets, &item.CurrentAssets, &item.Pledge)
		if err != nil {
			log.Printf("Error scanning GetMonitoringSummaryDetails row: %v", err)
			continue
		}
		monitoringsummary = append(monitoringsummary, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetMonitoringSummaryDetails: %v", err)
	}

	return monitoringsummary
}
