---
description: >-
  Overnight rates are interest rates that large banks use to borrow and lend
  from one another in the overnight market.
---

# Interbank Overnight Interest Rates

## Sterling Overnight Index Average \(SONIA\)

SONIA is is a risk-free rate based on actual transactions. It reflects the average of the interest rates that banks pay to borrow sterling overnight from other financial institutions. SONIA is the Working Group on Sterling Risk Free Reference Rates’ preferred benchmark for the transition to sterling risk-free rates from Libor. It is published every London business day by the Bank of England.

Source:  
[https://www.bankofengland.co.uk/markets/sonia-benchmark](https://www.bankofengland.co.uk/markets/sonia-benchmark)

Example for API call:  
[https://api.diadata.org/v1/interestrate/SONIA/2020-03-14](https://api.diadata.org/v1/interestrate/SONIA/2020-03-14)

## Secured Overnight Financing Rate \(SOFR\)

SOFR is a rate product produced by the Federal Reserve Bank of New York for the public good. It is a broad measure of the cost of borrowing cash overnight collateralized by Treasury securities.

More precisely, it is calculated as a volume-weighted median of transaction-level tri-party repo data collected from the Bank of New York Mellon as well as GCF Repo transaction data and data on bilateral Treasury repo transactions cleared through FICC's DVP service, which are obtained from DTCC Solutions LLC, an affiliate of the Depository Trust & Clearing Corporation. Each business day, the New York Fed publishes the SOFR on the New York Fed website at approximately 8:00 a.m. ET.

Sources:  
[https://apps.newyorkfed.org/markets/autorates/SO](https://apps.newyorkfed.org/markets/autorates/SOFR)[FR](https://apps.newyorkfed.org/markets/autorates/SOFR)  
[https://www.newyorkfed.org/medialibrary/Microsites/arrc/files/2019/Users\_Guide\_to\_SOFR.pdf](https://www.newyorkfed.org/medialibrary/Microsites/arrc/files/2019/Users_Guide_to_SOFR.pdf)

Example for API call:  
[https://api.diadata.org/v1/interestrate/SOFR/2020-04-16](https://api.diadata.org/v1/interestrate/SOFR/2020-04-16)

### SOFR Index \(SAFR\)

The SAFR measures the cumulative impact of compounding the SOFR on a unit of investment over time, with the initial value set to 1.00000000 on April 2, 2018, the first value date of the SOFR. The SAFR value reflects the effect of compounding the SOFR each business day and allows the calculation of compounded SOFR averages over custom time periods.

Each business day, the New York Fed publishes the SAFR on the New York Fed's website, shortly after the SOFR is published at approximately 8:00 a.m. ET.

Source:  
[https://apps.newyorkfed.org/markets/autorates/sofr-avg-ind\#Chart12](https://apps.newyorkfed.org/markets/autorates/sofr-avg-ind#Chart12)

Example for API call:  
[https://api.diadata.org/v1/interestrate/SAFR/2020-04-16](https://api.diadata.org/v1/interestrate/SAFR/2020-04-16)

### SOFR30, SOFR90 and SOFR180

These rates are compounded averages of the SOFR over rolling 30-, 90-, and 180-calendar day periods.

Source:  
[https://apps.newyorkfed.org/markets/autorates/sofr-avg-ind\#Chart12](https://apps.newyorkfed.org/markets/autorates/sofr-avg-ind#Chart12)

Example for API call:  
[https://api.diadata.org/v1/interestrate/SOFR90/2020-04-16](https://api.diadata.org/v1/interestrate/SOFR90/2020-04-16)

## Euro Short-Term Rate \(€STR\)

€STR is a reference rate for the currency euro. It is calculated by the European Central Bank \(ECB\) and is based on the money market statistical reporting of the Eurosystem. The euro short-term rate \(€STR\) is a rate which reflects the wholesale euro unsecured overnight borrowing costs of euro area banks.

It is calculated as a volume-weighted trimmed mean rounded to the third decimal. The volume-weighted trimmed mean is calculated by:

* ordering transactions from the lowest rate to the highest rate;
* aggregating the transactions occurring at each rate level;
* removing the top and bottom 25% in volume terms; and 
* calculating the mean of the remaining 50% of the volume-weighted distribution of rates.

A pro rata calculation is applied to volumes that span the thresholds for trimming to ensure that exactly 50% of the total eligible volume is used in the calculation of the volume-weighted mean.

The rate is published on every [TARGET2](https://en.wikipedia.org/wiki/TARGET2) business day at 8:00 CET \(reflecting the trading activity of the previous business day\). If errors are detected, the €STR is revised and republished on the same day at 9:00 CET.

Sources:  
[https://www.ecb.europa.eu/paym/initiatives/interest\_rate\_benchmarks/shared/pdf/ecb.ESTER\_methodology\_and\_policies.en.pdf](https://www.ecb.europa.eu/paym/initiatives/interest_rate_benchmarks/shared/pdf/ecb.ESTER_methodology_and_policies.en.pdf)  
[https://en.wikipedia.org/wiki/%E2%82%ACSTR](https://en.wikipedia.org/wiki/%E2%82%ACSTR)

**Remark:**  
In the DIA API we use the former abbreviation ESTER in order to avoid the special symbol €. Hence, an exemplary API call looks like:  
[https://api.diadata.org/v1/interestrate/ESTER/2020-04-16](https://api.diadata.org/v1/interestrate/ESTER/2020-04-16)

##  

