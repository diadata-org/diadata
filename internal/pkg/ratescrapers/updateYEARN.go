package ratescrapers

// UpdateYEARN retrieves value from smart contract and sends updated value through
// Channel s.chanInterestRate
func (s *RateScraper) UpdateYEARN() error {
	compoundAPRResult, err := s.yearnManager.GetAllCompoundAPR()
	if err != nil {
		return err
	}
	go s.yearnManager.Publish("APROracle", "CDAI", compoundAPRResult.CDAI, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CBAT", compoundAPRResult.CBAT, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CETH", compoundAPRResult.CETH, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CREP", compoundAPRResult.CREP, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CSAI", compoundAPRResult.CSAI, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CUSDC", compoundAPRResult.CUSDC, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CWBTC", compoundAPRResult.CWBTC, s.chanInterestRate)
	go s.yearnManager.Publish("APROracle", "CZRC", compoundAPRResult.CZRC, s.chanInterestRate)
	return nil
}

