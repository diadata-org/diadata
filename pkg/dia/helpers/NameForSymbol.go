package helpers

func SymbolIsName(symbol string) bool {
	switch symbol {
	case "STK":
		return true
	case "VITE":
		return true
	case "MEX":
		return true
	case "MCO":
		return true
	case "DADI":
		return true
	case "OST":
		return true
	case "SALT":
		return true
	case "QASH":
		return true
	case "EOS":
		return true
	case "NEO":
		return true
	case "OAX":
		return true
	case "BTF":
		return true
	case "NKN":
		return true
	case "VIBE":
		return true
	case "PIVX":
		return true
	case "IOST":
		return true
	case "WAX":
		return true
	case "YEE":
		return true
	case "CIC":
		return true
	case "BEC":
		return true
	case "TRA":
		return true
	case "OKB":
		return true
	case "HSR":
		return true
	case "CAI":
		return true
	default:
		return false
	}
}
func NameForSymbolManuallyAdded(symbol string) string {
	switch symbol {
	//GateIO pairs. Maybe inactive
	case "NKN":
		return "NKN"
	case "LLT":
		return "LLToken"
	case "BTF":
		return "BTF"
	case "PPS":
		return "Prophet"
	case "ADD":
		return "EOSADD"
	case "BXC":
		return "BonusCloud"
		// Bitfinex
	case "RRT":
		return "Recovery Right Token"
		//OKEx
	case "CIC":
		return "CIC"
	case "BEC":
		return "BEC"
	case "TRA":
		return "TRA"
	case "OKB":
		return "OKB"
	case "HSR":
		return "HSR"
	case "CAI":
		return "CAI"
		//Simex
	case "WCO":
		return "Winco"
	case "NBTK":
		return "Nebeus Crypto Bank"
	case "GOLD":
		return "Goldmint"
	case "MAD":
		return "MAD Network"
	case "ULTC":
		return "Ultimate Coin"
	case "BZS":
		return "Baista"
	case "RBM":
		return "Robomed"
	case "VOCT":
		return "VOCO"
	case "MTRC":
		return "ModulTrade"
	default:
		return symbol
	}
}
func NameForSymbol(symbol string) string {
	switch symbol {
	case "DIA":
		return "DIAData"
	case "BTC":
		return "Bitcoin"
	case "LTC":
		return "Litecoin"
	case "NMC":
		return "Namecoin"
	case "TRC":
		return "Terracoin"
	case "PPC":
		return "Peercoin"
	case "NVC":
		return "Novacoin"
	case "FTC":
		return "Feathercoin"
	case "MNC":
		return "Mincoin"
	case "FRC":
		return "Freicoin"
	case "IXC":
		return "Ixcoin"
	case "BTB":
		return "BitBar"
	case "WDC":
		return "WorldCoin"
	case "DGC":
		return "Digitalcoin"
	case "GLD":
		return "GoldCoin"
	case "ARG":
		return "Argentum"
	case "FST":
		return "Fastcoin"
	case "PXC":
		return "Phoenixcoin"
	case "MEC":
		return "Megacoin"
	case "IFC":
		return "Infinitecoin"
	case "XPM":
		return "Primecoin"
	case "ANC":
		return "Anoncoin"
	case "CSC":
		return "CasinoCoin"
	case "CBX":
		return "Bullion"
	case "EMD":
		return "Emerald Crypto"
	case "XRP":
		return "Ripple"
	case "QRK":
		return "Quark"
	case "ZET":
		return "Zetacoin"
	case "SRC":
		return "SecureCoin"
	case "SXC":
		return "Sexcoin"
	case "TAG":
		return "TagCoin"
	case "I0C":
		return "I0Coin"
	case "FLO":
		return "FLO"
	case "NXT":
		return "Nxt"
	case "UNO":
		return "Unobtanium"
	case "XJO":
		return "Joulecoin"
	case "DTC":
		return "Datacoin"
	case "BET":
		return "BetaCoin"
	case "GDC":
		return "GrandCoin"
	case "DEM":
		return "Deutsche eMark"
	case "DOGE":
		return "Dogecoin"
	case "NET":
		return "NetCoin"
	case "PHS":
		return "Philosopher Stones"
	case "DMD":
		return "Diamond"
	case "HBN":
		return "HoboNickels"
	case "TGC":
		return "Tigercoin"
	case "ORB":
		return "Orbitcoin"
	case "OMNI":
		return "Omni"
	case "CAT":
		return "Catcoin"
	case "TIPS":
		return "FedoraCoin"
	case "RPC":
		return "RonPaulCoin"
	case "MOON":
		return "Mooncoin"
	case "DIME":
		return "Dimecoin"
	case "42":
		return "42-coin"
	case "VTC":
		return "Vertcoin"
	case "KDC":
		return "KlondikeCoin"
	case "RED":
		return "RedCoin"
	case "DGB":
		return "DigiByte"
	case "SMC":
		return "SmartCoin"
	case "TES":
		return "TeslaCoin"
	case "NOBL":
		return "NobleCoin"
	case "RDD":
		return "ReddCoin"
	case "NYAN":
		return "Nyancoin"
	case "UTC":
		return "UltraCoin"
	case "POT":
		return "PotCoin"
	case "BLC":
		return "Blakecoin"
	case "MAX":
		return "MaxCoin"
	case "Q2C":
		return "QubitCoin"
	case "HUC":
		return "HunterCoin"
	case "DASH":
		return "Dash"
	case "XCP":
		return "Counterparty"
	case "CACH":
		return "CacheCoin"
	case "ICN":
		return "iCoin"
	case "MINT":
		return "MintCoin"
	case "ARI":
		return "Aricoin"
	case "DOPE":
		return "DopeCoin"
	case "AUR":
		return "Auroracoin"
	case "ANI":
		return "Animecoin"
	case "PTC":
		return "Pesetacoin"
	case "MARS":
		return "Marscoin"
	case "CASH":
		return "Cashcoin"
	case "PND":
		return "Pandacoin"
	case "MAZA":
		return "MAZA"
	case "UFO":
		return "Uniform Fiscal Object"
	case "BLK":
		return "BlackCoin"
	case "LTB":
		return "LiteBar"
	case "PHO":
		return "Photon"
	case "ZEIT":
		return "Zeitcoin"
	case "XMY":
		return "Myriad"
	case "NOTE":
		return "DNotes"
	case "SKC":
		return "Skeincoin"
	case "EMC2":
		return "Einsteinium"
	case "BTCS":
		return "Bitcoin Scrypt"
	case "CNO":
		return "Coin(O)"
	case "ECC":
		return "ECC"
	case "MONA":
		return "MonaCoin"
	case "RBY":
		return "Rubycoin"
	case "BELA":
		return "Bela"
	case "FLT":
		return "FlutterCoin"
	case "888":
		return "OctoCoin"
	case "FAIR":
		return "FairCoin"
	case "SLR":
		return "SolarCoin"
	case "EFL":
		return "e-Gulden"
	case "NLG":
		return "Gulden"
	case "PLC":
		return "Polcoin"
	case "GRS":
		return "Groestlcoin"
	case "XPD":
		return "PetroDollar"
	case "PLNC":
		return "PLNcoin"
	case "XWC":
		return "WhiteCoin"
	case "AC":
		return "AsiaCoin"
	case "POP":
		return "PopularCoin"
	case "BITS":
		return "Bitstar"
	case "QBC":
		return "Quebecoin"
	case "CCN":
		return "CannaCoin"
	case "BLU":
		return "BlueCoin"
	case "MAID":
		return "MaidSafeCoin"
	case "XBC":
		return "Bitcoin Plus"
	case "TALK":
		return "BTCtalkcoin"
	case "NYC":
		return "NewYorkCoin"
	case "CDN":
		return "Canada eCoin"
	case "GUN":
		return "Guncoin"
	case "PINK":
		return "PinkCoin"
	case "DRM":
		return "Dreamcoin"
	case "CFC":
		return "CoffeeCoin"
	case "ENRG":
		return "Energycoin"
	case "VRC":
		return "VeriCoin"
	case "TEK":
		return "TEKcoin"
	case "XMR":
		return "Monero"
	case "LCP":
		return "Litecoin Plus"
	case "CURE":
		return "Curecoin"
	case "UNB":
		return "UnbreakableCoin"
	case "CRYPT":
		return "CryptCoin"
	case "SUPER":
		return "SuperCoin"
	case "BOST":
		return "BoostCoin"
	case "HYPER":
		return "Hyper"
	case "BTQ":
		return "BitQuark"
	case "MOTO":
		return "Motocoin"
	case "CLOAK":
		return "CloakCoin"
	case "BSD":
		return "BitSend"
	case "C2":
		return "Coin2.1"
	case "FCN":
		return "Fantomcoin"
	case "BCN":
		return "Bytecoin"
	case "ABY":
		return "ArtByte"
	case "NAV":
		return "NavCoin"
	case "GRN":
		return "Granite"
	case "DON":
		return "Donationcoin"
	case "PIGGY":
		return "Piggycoin"
	case "START":
		return "Startcoin"
	case "KORE":
		return "Kore"
	case "XDN":
		return "DigitalNote"
	case "BBR":
		return "Boolberry"
	case "SHA":
		return "SHACoin"
	case "BLZ":
		return "BlazeCoin"
	case "THC":
		return "HempCoin"
	case "BRIT":
		return "BritCoin"
	case "XST":
		return "Stealth"
	case "TRUST":
		return "TrustPlus"
	case "CLAM":
		return "Clams"
	case "QTL":
		return "Quatloo"
	case "BTS":
		return "BitShares"
	case "TRK":
		return "Truckcoin"
	case "VIA":
		return "Viacoin"
	case "TRI":
		return "Triangles"
	case "IOC":
		return "I/O Coin"
	case "XCN":
		return "Cryptonite"
	case "CARBON":
		return "Carboncoin"
	case "CANN":
		return "CannabisCoin"
	case "XLM":
		return "Stellar"
	case "TIT":
		return "Titcoin"
	case "VTA":
		return "Virtacoin"
	case "HYP":
		return "HyperStake"
	case "J":
		return "Joincoin"
	case "SYS":
		return "Syscoin"
	case "BTM":
		return "Bytom"
	case "HAL":
		return "Halcyon"
	case "SJCX":
		return "Storjcoin X"
	case "NEOS":
		return "NeosCoin"
	case "EMC":
		return "Emercoin"
	case "RBBT":
		return "RabbitCoin"
	case "BURST":
		return "Burst"
	case "GAME":
		return "GameCredits"
	case "WSX":
		return "WeAreSatoshi"
	case "UBQ":
		return "Ubiq"
	case "BUN":
		return "BunnyCoin"
	case "OPAL":
		return "Opal"
	case "ACOIN":
		return "Acoin"
	case "FLDC":
		return "FoldingCoin"
	case "BITUSD":
		return "bitUSD"
	case "BITCNY":
		return "bitCNY"
	case "BITBTC":
		return "bitBTC"
	case "USNBT":
		return "NuBits"
	case "SLG":
		return "Sterlingcoin"
	case "XMG":
		return "Magi"
	case "EXCL":
		return "ExclusiveCoin"
	case "TROLL":
		return "Trollcoin"
	case "BSTY":
		return "GlobalBoost-Y"
	case "DP":
		return "DigitalPrice"
	case "PXI":
		return "Prime-XI"
	case "DSH":
		return "Dashcoin"
	case "AU":
		return "AurumCoin"
	case "STV":
		return "Sativacoin"
	case "XVG":
		return "Verge"
	case "NSR":
		return "NuShares"
	case "SPR":
		return "SpreadCoin"
	case "RBT":
		return "Rimbit"
	case "MUE":
		return "MonetaryUnit"
	case "BLOCK":
		return "Blocknet"
	case "GAP":
		return "Gapcoin"
	case "TTC":
		return "TittieCoin"
	case "CRW":
		return "Crown"
	case "BAY":
		return "BitBay"
	case "GCN":
		return "GCN Coin"
	case "XQN":
		return "Quotient"
	case "BYC":
		return "Bytecent"
	case "BCF":
		return "Bitcoin Fast"
	case "OK":
		return "OKCash"
	case "XPY":
		return "PayCoin"
	case "BITGOLD":
		return "bitGold"
	case "UIS":
		return "Unitus"
	case "GP":
		return "GoldPieces"
	case "COVAL":
		return "Circuits of Value"
	case "NXS":
		return "Nexus"
	case "SOON":
		return "SoonCoin"
	case "NKA":
		return "IncaKoin"
	case "SMLY":
		return "SmileyCoin"
	case "MAC":
		return "Machinecoin"
	case "BITSILVER":
		return "bitSilver"
	case "DOT":
		return "Dotcoin"
	case "KOBO":
		return "Kobocoin"
	case "BITB":
		return "Bean Cash"
	case "GEO":
		return "GeoCoin"
	case "USDT":
		return "Tether"
	case "WBB":
		return "Wild Beast Block"
	case "GRC":
		return "GridCoin"
	case "XCO":
		return "X-Coin"
	case "SAK":
		return "Sharkcoin"
	case "LDOGE":
		return "LiteDoge"
	case "SONG":
		return "SongCoin"
	case "LOG":
		return "Woodcoin"
	case "CRAVE":
		return "Crave"
	case "PURA":
		return "Pura"
	case "XEM":
		return "NEM"
	case "8BIT":
		return "8Bit"
	case "LEA":
		return "LeaCoin"
	case "NTRN":
		return "Neutron"
	case "XAUR":
		return "Xaurum"
	case "CF":
		return "Californium"
	case "AIB":
		return "Advanced Internet Blocks"
	case "SPHR":
		return "Sphere"
	case "MEDIC":
		return "MedicCoin"
	case "BUB":
		return "Bubble"
	case "XSD":
		return "SounDAC"
	case "UNIT":
		return "Universal Currency"
	case "PKB":
		return "ParkByte"
	case "ARB":
		return "ARbit"
	case "GAM":
		return "Gambit"
	case "BTA":
		return "Bata"
	case "ADC":
		return "AudioCoin"
	case "SNRG":
		return "Synergy"
	case "BITEUR":
		return "bitEUR"
	case "UNIC":
		return "UniCoin"
	case "FJC":
		return "FujiCoin"
	case "ERC":
		return "EuropeCoin"
	case "FUNK":
		return "The Cypherfunks"
	case "HXX":
		return "Hexx"
	case "XRA":
		return "Ratecoin"
	case "CREVA":
		return "CrevaCoin"
	case "IRL":
		return "IrishCoin"
	case "ZNY":
		return "Bitzeny"
	case "BSC":
		return "BowsCoin"
	case "ACP":
		return "AnarchistsPrime"
	case "CPN":
		return "CompuCoin"
	case "CHC":
		return "ChainCoin"
	case "SPRTS":
		return "Sprouts"
	case "HNC":
		return "Helleniccoin"
	case "CPC":
		return "Capricoin"
	case "FLAX":
		return "Flaxscript"
	case "MANNA":
		return "Manna"
	case "AXIOM":
		return "Axiom"
	case "LEO":
		return "LEOcoin"
	case "AEON":
		return "Aeon"
	case "ETH":
		return "Ethereum"
	case "SJW":
		return "SJWCoin"
	case "TX":
		return "TransferCoin"
	case "GCC":
		return "GuccioneCoin"
	case "AMS":
		return "AmsterdamCoin"
	case "EUC":
		return "Eurocoin"
	case "SC":
		return "Siacoin"
	case "GCR":
		return "Global Currency Reserve"
	case "SHIFT":
		return "Shift"
	case "VEC2":
		return "VectorAI"
	case "BOLI":
		return "Bolivarcoin"
	case "SPACE":
		return "SpaceCoin"
	case "BCY":
		return "Bitcrystals"
	case "PAK":
		return "Pakcoin"
	case "INFX":
		return "Influxcoin"
	case "EXP":
		return "Expanse"
	case "SIB":
		return "SIBCoin"
	case "SWING":
		return "Swing"
	case "FCT":
		return "Factom"
	case "DUO":
		return "ParallelCoin"
	case "SANDG":
		return "Save and Gain"
	case "PR":
		return "Prototanium"
	case "REP":
		return "Augur"
	case "SHND":
		return "StrongHands"
	case "$PAC":
		return "PACcoin"
	case "1337":
		return "Elite"
	case "$$$":
		return "Money"
	case "SOIL":
		return "SOILcoin"
	case "SCRT":
		return "SecretCoin"
	case "DFT":
		return "DraftCoin"
	case "OBITS":
		return "OBITS"
	case "AMP":
		return "Synereo"
	case "CLUB":
		return "ClubCoin"
	case "ADZ":
		return "Adzcoin"
	case "MOIN":
		return "Moin"
	case "AV":
		return "AvatarCoin"
	case "RC":
		return "RussiaCoin"
	case "EGC":
		return "EverGreenCoin"
	case "CRB":
		return "Creditbit"
	case "RADS":
		return "Radium"
	case "LTCR":
		return "Litecred"
	case "YOC":
		return "Yocoin"
	case "SLS":
		return "SaluS"
	case "FRN":
		return "Francs"
	case "EVIL":
		return "Evil Coin"
	case "DCR":
		return "Decred"
	case "PIVX":
		return "PIVX"
	case "SAFEX":
		return "Safe Exchange Coin"
	case "RBIES":
		return "Rubies"
	case "ADCN":
		return "Asiadigicoin"
	case "TRUMP":
		return "TrumpCoin"
	case "MEME":
		return "Memetic / PepeCoin"
	case "XCT":
		return "C-Bit"
	case "IMS":
		return "Independent Money System"
	case "HODL":
		return "HOdlcoin"
	case "BIGUP":
		return "BigUp"
	case "NEVA":
		return "NevaCoin"
	case "BUMBA":
		return "BumbaCoin"
	case "RVR":
		return "RevolutionVR"
	case "PEX":
		return "PosEx"
	case "CAB":
		return "Cabbage"
	case "MOJO":
		return "MojoCoin"
	case "GMX":
		return "GoldMaxCoin"
	case "LSK":
		return "Lisk"
	case "EDRC":
		return "EDRCoin"
	case "POST":
		return "PostCoin"
	case "BERN":
		return "BERNcash"
	case "QWARK":
		return "Qwark"
	case "DGD":
		return "DigixDAO"
	case "STEEM":
		return "Steem"
	case "FANS":
		return "Fantasy Cash"
	case "ESP":
		return "Espers"
	case "FUZZ":
		return "FuzzBalls"
	case "XHI":
		return "HiCoin"
	case "ARCO":
		return "AquariusCoin"
	case "XBTC21":
		return "Bitcoin 21"
	case "EL":
		return "Elcoin"
	case "ZUR":
		return "Zurcoin"
	case "611":
		return "SixEleven"
	case "2GIVE":
		return "2GIVE"
	case "XPTX":
		return "PlatinumBAR"
	case "ABC":
		return "Alphabit"
	case "LANA":
		return "LanaCoin"
	case "PONZI":
		return "PonziCoin"
	case "TESLA":
		return "TeslaCoilCoin"
	case "MXT":
		return "MarteXcoin"
	case "NLX":
		return "Nullex"
	case "RICHX":
		return "RichCoin"
	case "CTL":
		return "Citadel"
	case "WAVES":
		return "Waves"
	case "ICOO":
		return "ICO OpenLedger"
	case "PWR":
		return "PWR Coin"
	case "ION":
		return "ION"
	case "HVCO":
		return "High Voltage"
	case "XRE":
		return "RevolverCoin"
	case "GB":
		return "GoldBlocks"
	case "BRK":
		return "Breakout"
	case "DBTC":
		return "Debitcoin"
	case "CMT":
		return "Comet"
	case "RISE":
		return "Rise"
	case "CHESS":
		return "ChessCoin"
	case "LBC":
		return "LBRY Credits"
	case "PUT":
		return "PutinCoin"
	case "BRX":
		return "Breakout Stake"
	case "SYNX":
		return "Syndicate"
	case "CJ":
		return "Cryptojacks"
	case "HEAT":
		return "HEAT"
	case "SBD":
		return "Steem Dollars"
	case "ARDR":
		return "Ardor"
	case "ETC":
		return "Ethereum Classic"
	case "808":
		return "808Coin"
	case "BIT":
		return "First Bitcoin"
	case "ELE":
		return "Elementrem"
	case "KRB":
		return "Karbo"
	case "VPRC":
		return "VapersCoin"
	case "STRAT":
		return "Stratis"
	case "PRES":
		return "President Trump"
	case "ACES":
		return "Aces"
	case "GARY":
		return "President Johnson"
	case "TAJ":
		return "TajCoin"
	case "EDR":
		return "E-Dinar Coin"
	case "ATX":
		return "Artex Coin"
	case "XP":
		return "Experience Points"
	case "VLT":
		return "Veltor"
	case "KB3":
		return "B3Coin"
	case "GOLF":
		return "Golfcoin"
	case "NEO":
		return "NEO"
	case "LMC":
		return "LoMoCoin"
	case "BTDX":
		return "Bitcloud"
	case "NLC2":
		return "NoLimitCoin"
	case "VRM":
		return "VeriumReserve"
	case "ZYD":
		return "Zayedcoin"
	case "JIN":
		return "Jin Coin"
	case "XTO":
		return "Tao"
	case "PLU":
		return "Pluton"
	case "TELL":
		return "Tellurion"
	case "DLC":
		return "Dollarcoin"
	case "MST":
		return "MustangCoin"
	case "PROUD":
		return "PROUD Money"
	case "SEQ":
		return "Sequence"
	case "OMC":
		return "Omicron"
	case "1ST":
		return "FirstBlood"
	case "PEPECASH":
		return "Pepe Cash"
	case "SNGLS":
		return "SingularDTV"
	case "XZC":
		return "ZCoin"
	case "RCN":
		return "Rcoin"
	case "ATOM":
		return "Atomic Coin"
	case "JOBS":
		return "JobsCoin"
	case "TRIG":
		return "Triggers"
	case "SKR":
		return "Sakuracoin"
	case "LEVO":
		return "Levocoin"
	case "ARC":
		return "Advanced Technology Coin"
	case "QBT":
		return "Cubits"
	case "DMC":
		return "DynamicCoin"
	case "ZEC":
		return "Zcash"
	case "ASAFE2":
		return "AllSafe"
	case "BIP":
		return "BipCoin"
	case "ZCL":
		return "ZClassic"
	case "ZOI":
		return "Zoin"
	case "WA":
		return "WA Space"
	case "LKK":
		return "Lykke"
	case "GNT":
		return "Golem"
	case "ZMC":
		return "ZetaMicron"
	case "BTCR":
		return "Bitcurrency"
	case "IOP":
		return "Internet of People"
	case "VRS":
		return "Veros"
	case "HUSH":
		return "Hush"
	case "KURT":
		return "Kurrent"
	case "PASC":
		return "Pascal Coin"
	case "ENT":
		return "Eternity"
	case "INCNT":
		return "Incent"
	case "DCT":
		return "DECENT"
	case "GOLOS":
		return "Golos"
	case "NXC":
		return "Nexium"
	case "VSL":
		return "vSlice"
	case "DOLLAR":
		return "Dollar Online"
	case "VLTC":
		return "Vault Coin"
	case "PCS":
		return "Pabyosi Coin (Special)"
	case "GBYTE":
		return "Byteball Bytes"
	case "POSW":
		return "PoSW Coin"
	case "LUNA":
		return "Luna Coin"
	case "FRGC":
		return "Fargocoin"
	case "WINGS":
		return "Wings"
	case "DIX":
		return "Dix Asset"
	case "DAR":
		return "Darcrus"
	case "IFLT":
		return "InflationCoin"
	case "XSPEC":
		return "Spectrecoin"
	case "XSTC":
		return "Safe Trade Coin"
	case "BENJI":
		return "BenjiRolls"
	case "CCRB":
		return "CryptoCarbon"
	case "VIDZ":
		return "PureVidz"
	case "ICOB":
		return "ICOBID"
	case "IBANK":
		return "iBank"
	case "MKR":
		return "Maker"
	case "DRS":
		return "Digital Rupees"
	case "KMD":
		return "Komodo"
	case "FRST":
		return "FirstCoin"
	case "MGM":
		return "Magnum"
	case "TSE":
		return "Tattoocoin (Standard Edition)"
	case "SFC":
		return "Solarflarecoin"
	case "WCT":
		return "Waves Community Token"
	case "ICON":
		return "Iconic"
	case "KUSH":
		return "KushCoin"
	case "BOAT":
		return "BOAT"
	case "ERY":
		return "Eryllium"
	case "ELS":
		return "Elysium"
	case "GBG":
		return "Golos Gold"
	case "CNT":
		return "Centurion"
	case "MAR":
		return "Marijuanacoin"
	case "MSCN":
		return "Master Swiscoin"
	case "MLN":
		return "Melon"
	case "ALL":
		return "Allion"
	case "PRC":
		return "PRCoin"
	case "TIME":
		return "Chronobank"
	case "ARGUS":
		return "Argus"
	case "RNS":
		return "Renos"
	case "SWT":
		return "Swarm City"
	case "PIE":
		return "PIECoin"
	case "MARX":
		return "MarxCoin"
	case "VISIO":
		return "Visio"
	case "NANO":
		return "Nano"
	case "GEERT":
		return "GeertCoin"
	case "PASL":
		return "Pascal Lite"
	case "MILO":
		return "MiloCoin"
	case "MUSIC":
		return "Musicoin"
	case "ZER":
		return "Zero"
	case "HONEY":
		return "Honey"
	case "NETKO":
		return "Netko"
	case "ARK":
		return "Ark"
	case "DYN":
		return "Dynamic"
	case "TKS":
		return "Tokes"
	case "MER":
		return "Mercury"
	case "TAAS":
		return "TaaS"
	case "SOAR":
		return "Soarcoin"
	case "EDG":
		return "Edgeless"
	case "B@":
		return "Bankcoin"
	case "ZSE":
		return "ZSEcoin"
	case "WORM":
		return "HealthyWormCoin"
	case "DTB":
		return "Databits"
	case "UNI":
		return "Universe"
	case "XLR":
		return "Solaris"
	case "IMX":
		return "Impact"
	case "XAS":
		return "Asch"
	case "DBIX":
		return "DubaiCoin"
	case "KED":
		return "Darsek"
	case "GUP":
		return "Matchpool"
	case "USC":
		return "Ultimate Secure Cash"
	case "ECN":
		return "E-coin"
	case "SKY":
		return "Skycoin"
	case "BLAZR":
		return "BlazerCoin"
	case "ATMOS":
		return "Atmos"
	case "HPC":
		return "Happycoin"
	case "ZENI":
		return "Zennies"
	case "CXT":
		return "Coinonat"
	case "XOT":
		return "Internet of Things"
	case "CONX":
		return "Concoin"
	case "XBY":
		return "XTRABYTES"
	case "RLC":
		return "iExec RLC"
	case "TRST":
		return "WeTrust"
	case "DEUS":
		return "DeusCoin"
	case "ALT":
		return "Altcoin"
	case "WGO":
		return "WavesGo"
	case "MCRN":
		return "MACRON"
	case "TLE":
		return "Tattoocoin (Limited Edition)"
	case "PROC":
		return "ProCurrency"
	case "SCS":
		return "SpeedCash"
	case "BTX":
		return "Bitcore"
	case "VOLT":
		return "Bitvolt"
	case "LUN":
		return "Lunyr"
	case "GNO":
		return "Gnosis"
	case "TKN":
		return "TokenCard"
	case "RAIN":
		return "Condensate"
	case "GPL":
		return "Gold Pressed Latinum"
	case "HMQ":
		return "Humaniq"
	case "ITI":
		return "iTicoin"
	case "MNE":
		return "Minereum"
	case "CNNC":
		return "Cannation"
	case "CREA":
		return "Creativecoin"
	case "DICE":
		return "Etheroll"
	case "INSN":
		return "InsaneCoin"
	case "HALLO":
		return "Halloween Coin"
	case "ANT":
		return "Aragon"
	case "PZM":
		return "PRIZM"
	case "RLT":
		return "RouletteToken"
	case "QTUM":
		return "Qtum"
	case "ECO":
		return "EcoCoin"
	case "EQT":
		return "EquiTrader"
	case "DMB":
		return "Digital Money Bits"
	case "APX":
		return "APX"
	case "MCAP":
		return "MCAP"
	case "NANOX":
		return "Project-X"
	case "MAY":
		return "Theresa May Coin"
	case "SUMO":
		return "Sumokoin"
	case "ZENGOLD":
		return "ZenGold"
	case "BAT":
		return "Basic Attention Token"
	case "ZEN":
		return "Horizen"
	case "ETBS":
		return "Ethbits"
	case "AE":
		return "Aeternity"
	case "V":
		return "Version"
	case "ETP":
		return "Metaverse ETP"
	case "EBST":
		return "eBoost"
	case "ADK":
		return "Aidos Kuneen"
	case "STEX":
		return "STEX"
	case "PTOY":
		return "Patientory"
	case "VERI":
		return "Veritaseum"
	case "ECA":
		return "Electra"
	case "QRL":
		return "Quantum Resistant Ledger"
	case "ETT":
		return "EncryptoTel WAVES"
	case "MGO":
		return "MobileGo"
	case "AMMO":
		return "Ammo Reloaded"
	case "NRO":
		return "Neuro"
	case "PPY":
		return "Peerplays"
	case "MIOTA":
		return "IOTA"
	case "MYST":
		return "Mysterium"
	case "MORE":
		return "More Coin"
	case "SNM":
		return "SONM"
	case "LINX":
		return "Linx"
	case "ZRC":
		return "ZrCoin"
	case "BNT":
		return "Bancor"
	case "CHEAP":
		return "Cheapcoin"
	case "CFI":
		return "Cofound.it"
	case "GLT":
		return "GlobalToken"
	case "NMR":
		return "Numeraire"
	case "UNIFY":
		return "Unify"
	case "XEL":
		return "Elastic"
	case "COUPE":
		return "Coupecoin"
	case "MRT":
		return "Miners' Reward Token"
	case "BITOK":
		return "Bitok"
	case "KNC":
		return "KingN Coin"
	case "DCY":
		return "Dinastycoin"
	case "XLC":
		return "Leviar"
	case "ONX":
		return "Onix"
	case "BTPL":
		return "Bitcoin Planet"
	case "GXS":
		return "GXChain"
	case "ATCC":
		return "ATC Coin"
	case "GOOD":
		return "Goodomy"
	case "ANTX":
		return "Antimatter"
	case "BRO":
		return "Bitradio"
	case "FLASH":
		return "Flash"
	case "FUN":
		return "FunFair"
	case "PAY":
		return "TenX"
	case "SNT":
		return "Status"
	case "CHAN":
		return "ChanCoin"
	case "EFYT":
		return "Ergo"
	case "BRIA":
		return "BriaCoin"
	case "EOS":
		return "EOS"
	case "TURBO":
		return "TurboCoin"
	case "ADX":
		return "AdEx"
	case "DNR":
		return "Denarius"
	case "STORJ":
		return "Storj"
	case "BNX":
		return "BnrtxCoin"
	case "SOCC":
		return "SocialCoin"
	case "ADT":
		return "adToken"
	case "MCO":
		return "Crypto.com"
	case "PING":
		return "CryptoPing"
	case "UET":
		return "Useless Ethereum Token"
	case "WGR":
		return "Wagerr"
	case "SLEVIN":
		return "Slevin"
	case "ECOB":
		return "Ecobit"
	case "UNRC":
		return "UniversalRoyalCoin"
	case "PLBT":
		return "Polybius"
	case "GAS":
		return "Gas"
	case "SNC":
		return "SunContract"
	case "JET":
		return "Jetcoin"
	case "MTL":
		return "Metal"
	case "PPT":
		return "Populous"
	case "WOMEN":
		return "WomenCoin"
	case "BDL":
		return "Bitdeal"
	case "XID":
		return "Sphre AIR "
	case "DAXX":
		return "DaxxCoin"
	case "RUP":
		return "Rupee"
	case "PCN":
		return "PeepCoin"
	case "HERO":
		return "Sovereign Hero"
	case "SAN":
		return "Santiment Network Token"
	case "OMG":
		return "OmiseGO"
	case "TER":
		return "TerraNova"
	case "CVN":
		return "CVCoin"
	case "XRL":
		return "Rialto"
	case "LINDA":
		return "Linda"
	case "MBRS":
		return "Embers"
	case "CVC":
		return "Civic"
	case "ETHOS":
		return "Ethos"
	case "BTWTY":
		return "Bit20"
	case "STA":
		return "Starta"
	case "COAL":
		return "BitCoal"
	case "LBTC":
		return "LiteBitcoin"
	case "PART":
		return "Particl"
	case "SMART":
		return "SmartCash"
	case "SKIN":
		return "SkinCoin"
	case "BCH":
		return "Bitcoin Cash"
	case "HMC":
		return "HarmonyCoin"
	case "TOA":
		return "ToaCoin"
	case "PLR":
		return "Pillar"
	case "XRY":
		return "Royalties"
	case "SIGT":
		return "Signatum"
	case "CTIC2":
		return "Coimatic 2.0"
	case "OCT":
		return "OracleChain"
	case "BNB":
		return "Binance Coin"
	case "300":
		return "300 Token"
	case "PBT":
		return "Primalbase Token"
	case "CMPCO":
		return "CampusCoin"
	case "EMB":
		return "EmberCoin"
	case "IXT":
		return "IXT"
	case "GSR":
		return "GeyserCoin"
	case "MSP":
		return "Mothership"
	case "BIRDS":
		return "Birds"
	case "CRM":
		return "Cream"
	case "ERA":
		return "ERA"
	case "KEK":
		return "KekCoin"
	case "OAX":
		return "OAX"
	case "DNT":
		return "district0x"
	case "FYN":
		return "FundYourselfNow"
	case "STX":
		return "Stox"
	case "MINEX":
		return "Minex"
	case "CDT":
		return "Blox"
	case "WINK":
		return "Wink"
	case "MAO":
		return "Mao Zedong"
	case "BITCF":
		return "First Bitcoin Capital"
	case "NDC":
		return "NEVERDIE"
	case "TIX":
		return "Blocktix"
	case "DCN":
		return "Dentacoin"
	case "RUPX":
		return "Rupaya"
	case "SHDW":
		return "Shadow Token"
	case "ONION":
		return "DeepOnion"
	case "ADST":
		return "AdShares"
	case "DDF":
		return "DigitalDevelopersFund"
	case "BAS":
		return "BitAsean"
	case "DENT":
		return "Dent"
	case "MBI":
		return "Monster Byte"
	case "IFT":
		return "InvestFeed"
	case "XCXT":
		return "CoinonatX"
	case "RIYA":
		return "Etheriya"
	case "TCC":
		return "The ChampCoin"
	case "ZRX":
		return "0x"
	case "BLN":
		return "Bolenum"
	case "SMOKE":
		return "Smoke"
	case "YOYOW":
		return "YOYOW"
	case "GRWI":
		return "Growers International"
	case "MYB":
		return "MyBit"
	case "HC":
		return "HyperCash"
	case "TFL":
		return "TrueFlip"
	case "NAS":
		return "Nebulas"
	case "DALC":
		return "Dalecoin"
	case "PRN":
		return "Protean"
	case "ACC":
		return "AdCoin"
	case "BBP":
		return "BiblePay"
	case "BQ":
		return "bitqy"
	case "ACT":
		return "Achain"
	case "NAMO":
		return "NamoCoin"
	case "SIGMA":
		return "SIGMAcoin"
	case "XMCC":
		return "Monoeci"
	case "TNT":
		return "Tierion"
	case "WTC":
		return "Waltonchain"
	case "BRAT":
		return "BROTHER"
	case "PST":
		return "Primas"
	case "OPT":
		return "Opus"
	case "SUR":
		return "Suretly"
	case "LRC":
		return "Loopring"
	case "LTCU":
		return "LiteCoin Ultra"
	case "POE":
		return "Po.et"
	case "NTO":
		return "Fujinto"
	case "STRC":
		return "StarCredits"
	case "KRONE":
		return "Kronecoin"
	case "MCC":
		return "Moving Cloud Coin"
	case "CYDER":
		return "Cyder"
	case "MTNC":
		return "Masternodecoin"
	case "MTH":
		return "Monetha"
	case "AVT":
		return "Aventus"
	case "DLT":
		return "Agrello"
	case "HVN":
		return "Hiveterminal Token"
	case "VSX":
		return "Vsync"
	case "MAGN":
		return "Magnetcoin"
	case "MDA":
		return "Moeda Loyalty Points"
	case "NEBL":
		return "Neblio"
	case "VIVO":
		return "VIVO"
	case "TRX":
		return "TRON"
	case "OCL":
		return "Oceanlab"
	case "REX":
		return "imbrex"
	case "BUZZ":
		return "BuzzCoin"
	case "CREDO":
		return "Credo"
	case "DRXNE":
		return "DROXNE"
	case "AHT":
		return "Bowhead"
	case "MANA":
		return "Decentraland"
	case "IND":
		return "Indorse Token"
	case "XPA":
		return "XPA"
	case "SCL":
		return "Sociall"
	case "ATB":
		return "ATBCoin"
	case "IQT":
		return "iQuant"
	case "ETHD":
		return "Ethereum Dark"
	case "PRO":
		return "Propy"
	case "LINK":
		return "Chainlink"
	case "BMC":
		return "Blackmoon"
	case "WIC":
		return "Wi Coin"
	case "ELIX":
		return "Elixir"
	case "XBL":
		return "Billionaire Token"
	case "VIBE":
		return "VIBE"
	case "SUB":
		return "Substratum"
	case "DAY":
		return "Chronologic"
	case "CHIPS":
		return "CHIPS"
	case "TKR":
		return "CryptoInsight"
	case "PIX":
		return "Lampix"
	case "COSS":
		return "COSS"
	case "CSNO":
		return "BitDice"
	case "RVT":
		return "Rivetz"
	case "KIN":
		return "Kin"
	case "ITZ":
		return "Interzone"
	case "TGT":
		return "Target Coin"
	case "SALT":
		return "SALT"
	case "ORME":
		return "Ormeus Coin"
	case "KLN":
		return "Kolion"
	case "MCI":
		return "Musiconomi"
	case "COLX":
		return "ColossusXT"
	case "TZC":
		return "TrezarCoin"
	case "HDLB":
		return "HODL Bucks"
	case "ODN":
		return "Obsidian"
	case "COB":
		return "Cobinhood"
	case "REC":
		return "Regalcoin"
	case "MSD":
		return "MSD"
	case "BIS":
		return "Bismuth"
	case "ADA":
		return "Cardano"
	case "XTZ":
		return "Tezos"
	case "VOISE":
		return "Voise"
	case "XIN":
		return "Infinity Economics"
	case "ATM":
		return "ATMChain"
	case "KICK":
		return "KickCoin"
	case "VIB":
		return "Viberate"
	case "RHOC":
		return "RChain"
	case "INXT":
		return "Internxt"
	case "WHL":
		return "WhaleCoin"
	case "FLIK":
		return "FLiK"
	case "EBET":
		return "EthBet"
	case "CNX":
		return "Cryptonex"
	case "WILD":
		return "Wild Crypto"
	case "REAL":
		return "REAL"
	case "HBT":
		return "Hubii Network"
	case "CCT":
		return "Crystal Clear "
	case "BCO":
		return "BridgeCoin"
	case "EVX":
		return "Everex"
	case "PPP":
		return "PayPie"
	case "AIR":
		return "AirToken"
	case "POS":
		return "PoSToken"
	case "SDRN":
		return "Senderon"
	case "ALIS":
		return "ALIS"
	case "BTCZ":
		return "BitcoinZ"
	case "HGT":
		return "HelloGold"
	case "CND":
		return "Cindicator"
	case "ENG":
		return "Enigma"
	case "CTIC3":
		return "Coimatic 3.0"
	case "BSN":
		return "Bastonet"
	case "ZSC":
		return "Zeusshield"
	case "ECASH":
		return "Ethereum Cash"
	case "COR":
		return "CORION"
	case "SIC":
		return "Swisscoin"
	case "ATS":
		return "Authorship"
	case "RKC":
		return "Royal Kingdom Coin"
	case "AKY":
		return "Akuya Coin"
	case "EXN":
		return "ExchangeN"
	case "PIPL":
		return "PiplCoin"
	case "EDO":
		return "Eidoo"
	case "AST":
		return "AirSwap"
	case "BSR":
		return "BitSoar"
	case "CAG":
		return "Change"
	case "BCPT":
		return "BlockMason Credit Protocol"
	case "AION":
		return "Aion"
	case "TRCT":
		return "Tracto"
	case "ART":
		return "Maecenas"
	case "XGOX":
		return "XGOX"
	case "EVR":
		return "Everus"
	case "DUTCH":
		return "Dutch Coin"
	case "SWP":
		return "Swapcoin"
	case "OTN":
		return "Open Trading Network"
	case "DRT":
		return "DomRaider"
	case "REQ":
		return "Request Network"
	case "B2X":
		return "SegWit2x"
	case "ETG":
		return "Ethereum Gold"
	case "BLUE":
		return "Blue Protocol"
	case "RUNNERS":
		return "Runners"
	case "LIFE":
		return "LIFE"
	case "HDG":
		return "Hedge"
	case "MOD":
		return "Modum"
	case "AMB":
		return "Ambrosus"
	case "ICOS":
		return "ICOS"
	case "BTG":
		return "Bitcoin Gold"
	case "KCS":
		return "KuCoin Shares"
	case "EXRN":
		return "EXRNchain"
	case "POLL":
		return "ClearPoll"
	case "LA":
		return "LATOKEN"
	case "XUC":
		return "Exchange Union"
	case "NULS":
		return "Nuls"
	case "BTCRED":
		return "Bitcoin Red"
	case "PRG":
		return "Paragon"
	case "BOS":
		return "BOScoin"
	case "GMT":
		return "Mercury Protocol"
	case "ICX":
		return "ICON"
	case "JS":
		return "JavaScript Token"
	case "ELITE":
		return "Ethereum Lite"
	case "ITT":
		return "Intelligent Trading Foundation"
	case "IETH":
		return "iEthereum"
	case "PIRL":
		return "Pirl"
	case "XNN":
		return "Xenon"
	case "LUX":
		return "LUXCoin"
	case "NTWK":
		return "Network Token"
	case "DOV":
		return "Dovu"
	case "PHX":
		return "Red Pulse Phoenix"
	case "BT2":
		return "BT2 CST"
	case "PLACO":
		return "PlayerCoin"
	case "ROOFS":
		return "Roofs"
	case "FAP":
		return "FAPcoin"
	case "BTCM":
		return "BTCMoon"
	case "FUEL":
		return "Etherparty"
	case "ELLA":
		return "Ellaism"
	case "QVT":
		return "Qvolta"
	case "RMC":
		return "Russian Miner Coin"
	case "FYP":
		return "FlypMe"
	case "EBTC":
		return "eBitcoin"
	case "BTBc":
		return "Bitbase"
	case "ENJ":
		return "Enjin Coin"
	case "IBTC":
		return "iBTC"
	case "POWR":
		return "Power Ledger"
	case "GRID":
		return "Grid+"
	case "R":
		return "Revain"
	case "ATL":
		return "ATLANT"
	case "ETN":
		return "Electroneum"
	case "HIGH":
		return "High Gain"
	case "MNX":
		return "MinexCoin"
	case "SONO":
		return "SONO"
	case "FOR":
		return "FORCE"
	case "DATA":
		return "Streamr DATAcoin"
	case "XSH":
		return "SHIELD"
	case "ELTCOIN":
		return "ELTCOIN"
	case "DSR":
		return "Desire"
	case "UKG":
		return "Unikoin Gold"
	case "CRDNC":
		return "Credence Coin"
	case "NIO":
		return "Autonio"
	case "CTX":
		return "CarTaxi Token"
	case "ARN":
		return "Aeron"
	case "STARS":
		return "StarCash Network"
	case "PHR":
		return "Phore"
	case "INN":
		return "Innova"
	case "RDN":
		return "Raiden Network Token"
	case "DPY":
		return "Delphy"
	case "ZEPH":
		return "Zephyr"
	case "ERC20":
		return "ERC20"
	case "TIE":
		return "Ties.DB"
	case "BPL":
		return "Blockpool"
	case "GRIM":
		return "Grimcoin"
	case "EPY":
		return "Emphy"
	case "NEOG":
		return "NEO GOLD"
	case "DBET":
		return "DecentBet"
	case "HST":
		return "Decision Token"
	case "SSS":
		return "Sharechain"
	case "UFR":
		return "Upfiring"
	case "STU":
		return "bitJob"
	case "GVT":
		return "Genesis Vision"
	case "EAGLE":
		return "EagleCoin"
	case "EAG":
		return "EA Coin"
	case "PRIX":
		return "Privatix"
	case "LTHN":
		return "Lethean"
	case "EBCH":
		return "EBCH"
	case "ASTRO":
		return "Astro"
	case "PAYX":
		return "Paypex"
	case "GRX":
		return "GOLD Reward Token"
	case "BTE":
		return "BitSerial"
	case "SGR":
		return "Sugar Exchange"
	case "VIU":
		return "Viuly"
	case "XLQ":
		return "ALQO"
	case "GBX":
		return "GoByte"
	case "WC":
		return "WINCOIN"
	case "PRL":
		return "Oyster"
	case "B2B":
		return "B2BX"
	case "PNX":
		return "Phantomx"
	case "DGPT":
		return "DigiPulse"
	case "DNA":
		return "EncrypGen"
	case "INK":
		return "Ink"
	case "BOT":
		return "Bodhi"
	case "QSP":
		return "Quantstamp"
	case "QASH":
		return "QASH"
	case "ZZC":
		return "ZoZoCoin"
	case "TSL":
		return "Energo"
	case "PBL":
		return "Publica"
	case "MAG":
		return "Magnet"
	case "SPANK":
		return "SpankChain"
	case "VOT":
		return "VoteCoin"
	case "BCD":
		return "Bitcoin Diamond"
	case "VEE":
		return "BLOCKv"
	case "AI":
		return "POLY AI"
	case "PLX":
		return "PlexCoin"
	case "DIVX":
		return "Divi Exchange Token"
	case "MONK":
		return "Monkey Project"
	case "FLIXX":
		return "Flixxo"
	case "GLS":
		return "GlassCoin"
	case "TNB":
		return "Time New Bank"
	case "WISH":
		return "MyWish"
	case "EVC":
		return "EventChain"
	case "LEND":
		return "ETHLend"
	case "ONG":
		return "SoMee.Social"
	case "CCO":
		return "Ccore"
	case "DRGN":
		return "Dragonchain"
	case "PFR":
		return "Payfair"
	case "PRE":
		return "Presearch"
	case "BCDN":
		return "BlockCDN"
	case "CAPP":
		return "Cappasity"
	case "ERO":
		return "Eroscoin"
	case "MAGE":
		return "MagicCoin"
	case "ITC":
		return "IoT Chain"
	case "JIYO":
		return "Jiyo OLD"
	case "SEND":
		return "Social Send"
	case "BON":
		return "Bonpay"
	case "NUKO":
		return "Nekonium"
	case "SNOV":
		return "Snovian.Space"
	case "BWK":
		return "Bulwark"
	case "SAGA":
		return "SagaCoin"
	case "CMS":
		return "COMSA ETH"
	case "KBR":
		return "Kubera Coin"
	case "TOK":
		return "Tokugawa"
	case "PCOIN":
		return "Pioneer Coin"
	case "WABI":
		return "WaBi"
	case "CRC":
		return "CrowdCoin"
	case "WAND":
		return "WandX"
	case "SPF":
		return "SportyCo"
	case "CRED":
		return "Verify"
	case "SCT":
		return "Soma"
	case "UQC":
		return "Uquid Coin"
	case "MDS":
		return "MediShares"
	case "PRA":
		return "ProChain"
	case "IGNIS":
		return "Ignis"
	case "SMT":
		return "SmartMesh"
	case "HWC":
		return "HollyWoodCoin"
	case "PKT":
		return "Playkey"
	case "FIL":
		return "Filecoin Futures"
	case "BCX":
		return "BitcoinX"
	case "SBTC":
		return "Super Bitcoin"
	case "DAT":
		return "Datum"
	case "TRDT":
		return "Trident Group"
	case "AMM":
		return "MicroMoney"
	case "LOC":
		return "LockTrip"
	case "WRC":
		return "Worldcore"
	case "GTO":
		return "Gifto"
	case "YTN":
		return "YENTEN"
	case "GNX":
		return "Genaro Network"
	case "UBTC":
		return "United Bitcoin"
	case "STAR":
		return "Starbase"
	case "OST":
		return "OST"
	case "STORM":
		return "Storm"
	case "DTR":
		return "Dynamic Trading Rights"
	case "ELF":
		return "aelf"
	case "WAX":
		return "WAX"
	case "MED":
		return "MediBloc QRC20"
	case "DEW":
		return "DEW"
	case "NGC":
		return "NAGA"
	case "BRD":
		return "Bread"
	case "BIX":
		return "Bibox Token"
	case "DAI":
		return "Dai"
	case "SPHTX":
		return "SophiaTX"
	case "BNTY":
		return "Bounty0x"
	case "ACE":
		return "ACE (TokenStars)"
	case "DIM":
		return "DIMCOIN"
	case "SRN":
		return "SIRIN LABS Token"
	case "CPAY":
		return "Cryptopay"
	case "HTML":
		return "HTMLCOIN"
	case "DBC":
		return "DeepBrain Chain"
	case "HBC":
		return "HomeBlockCoin"
	case "NEU":
		return "Neumark"
	case "UTK":
		return "UTRUST"
	case "QLC":
		return "QLC Chain"
	case "PLAY":
		return "HEROcoin"
	case "MTX":
		return "Matryx"
	case "MOT":
		return "Olympus Labs"
	case "HPY":
		return "Hyper Pay"
	case "PYLNT":
		return "Pylon Network"
	case "STAK":
		return "STRAKS"
	case "FDX":
		return "FidentiaX"
	case "GTC":
		return "Game.com"
	case "TAU":
		return "Lamden"
	case "BLT":
		return "Bloom"
	case "SWFTC":
		return "SwftCoin"
	case "COV":
		return "Covesting"
	case "CAN":
		return "CanYaCoin"
	case "APPC":
		return "AppCoins"
	case "HPB":
		return "High Performance Blockchain"
	case "WICC":
		return "WaykiChain"
	case "MDT":
		return "Measurable Data Token"
	case "GCS":
		return "GameChain System"
	case "NMS":
		return "Numus"
	case "CL":
		return "Coinlancer"
	case "CEFS":
		return "CryptopiaFeeShares"
	case "GET":
		return "GET Protocol"
	case "OPC":
		return "OP Coin"
	case "CFUN":
		return "CFun"
	case "AIDOC":
		return "AI Doctor"
	case "POLIS":
		return "Polis"
	case "HKN":
		return "Hacken"
	case "SHOW":
		return "Show"
	case "STN":
		return "Steneum Coin"
	case "ZAP":
		return "Zap"
	case "TCT":
		return "TokenClub"
	case "AIX":
		return "Aigang"
	case "REBL":
		return "REBL"
	case "INS":
		return "INS Ecosystem"
	case "GOD":
		return "Bitcoin God"
	case "UTT":
		return "United Traders Token"
	case "CDX":
		return "Commodity Ad Network"
	case "TIO":
		return "Trade Token"
	case "BDG":
		return "BitDegree"
	case "QUN":
		return "QunQun"
	case "TOPC":
		return "TopChain"
	case "LEV":
		return "Leverj"
	case "KRM":
		return "Karma"
	case "KCASH":
		return "Kcash"
	case "ATN":
		return "ATN"
	case "SXDT":
		return "Spectre.ai Dividend Token"
	case "SXUT":
		return "Spectre.ai Utility Token"
	case "SWTC":
		return "Jingtum Tech"
	case "VZT":
		return "Vezt"
	case "CLD":
		return "Cloud"
	case "UCOM":
		return "UCOM"
	case "BCA":
		return "Bitcoin Atom"
	case "UGC":
		return "ugChain"
	case "BKX":
		return "Bankex"
	case "EKO":
		return "EchoLink"
	case "BTO":
		return "Bottos"
	case "AT":
		return "AWARE"
	case "TEL":
		return "Telcoin"
	case "IC":
		return "Ignition"
	case "WETH":
		return "WETH"
	case "HAC":
		return "Hackspace Capital"
	case "KEY":
		return "Selfkey"
	case "INT":
		return "Internet Node Token"
	case "RNT":
		return "OneRoot Network"
	case "SENSE":
		return "Sense"
	case "MOAC":
		return "MOAC"
	case "TOKC":
		return "TOKYO"
	case "IOST":
		return "IOST"
	case "IDT":
		return "InvestDigital"
	case "AIT":
		return "AICHAIN"
	case "QUBE":
		return "Qube"
	case "EDT":
		return "EtherDelta Token"
	case "SPC":
		return "SpaceChain"
	case "ORE":
		return "Galactrum"
	case "HORSE":
		return "Ethorse"
	case "RCT":
		return "RealChain"
	case "ARCT":
		return "ArbitrageCT"
	case "THETA":
		return "Theta Token"
	case "MVC":
		return "Maverick Chain"
	case "NOX":
		return "Nitro"
	case "IPL":
		return "VouchForMe"
	case "IDXM":
		return "IDEX Membership"
	case "AURA":
		return "Aurora DAO"
	case "AGI":
		return "SingularityNET"
	case "GAT":
		return "Gatcoin"
	case "SEXC":
		return "ShareX"
	case "CHAT":
		return "ChatCoin"
	case "DDD":
		return "Scry.info"
	case "MOBI":
		return "Mobius"
	case "HOT":
		return "Hydro Protocol"
	case "STC":
		return "StarChain"
	case "IPC":
		return "IPChain"
	case "LIGHT":
		return "LightChain"
	case "REF":
		return "RefToken"
	case "YEE":
		return "YEE"
	case "AAC":
		return "Acute Angle Cloud"
	case "SSC":
		return "SelfSell"
	case "READ":
		return "Read"
	case "MOF":
		return "Molecular Future"
	case "TNC":
		return "Trinity Network Credit"
	case "C20":
		return "CRYPTO20"
	case "ARY":
		return "Block Array"
	case "DTA":
		return "DATA"
	case "CRPT":
		return "Crypterium"
	case "SPK":
		return "Sparks"
	case "CV":
		return "carVertical"
	case "TBX":
		return "Tokenbox"
	case "EKT":
		return "EDUCare"
	case "UIP":
		return "UnlimitedIP"
	case "PRS":
		return "PressOne"
	case "OF":
		return "OFCOIN"
	case "TRUE":
		return "TrueChain"
	case "OCN":
		return "Odyssey"
	case "IDH":
		return "indaHash"
	case "QBIC":
		return "Qbic"
	case "GUESS":
		return "Peerguess"
	case "AID":
		return "AidCoin"
	case "EVE":
		return "Devery"
	case "BPT":
		return "Blockport"
	case "AXPR":
		return "aXpire"
	case "TRAC":
		return "OriginTrail"
	case "LET":
		return "LinkEye"
	case "ZIL":
		return "Zilliqa"
	case "MEET":
		return "CoinMeet"
	case "SLT":
		return "Smartlands"
	case "FOTA":
		return "Fortuna"
	case "SOC":
		return "All Sports"
	case "MAN":
		return "Matrix AI Network"
	case "GRLC":
		return "Garlicoin"
	case "RUFF":
		return "Ruff"
	case "NKC":
		return "Nework"
	case "COFI":
		return "CoinFi"
	case "EQL":
		return "Equal"
	case "HLC":
		return "HalalChain"
	case "ZPT":
		return "Zeepin"
	case "OC":
		return "OceanChain"
	case "CANDY":
		return "Candy"
	case "SMS":
		return "Speed Mining Service"
	case "EPC":
		return "Electronic PK Chain"
	case "VLC":
		return "ValueChain"
	case "BTW":
		return "BitWhite"
	case "CXO":
		return "CargoX"
	case "TRF":
		return "Travelflex"
	case "ELA":
		return "Elastos"
	case "STK":
		return "STK"
	case "PARETO":
		return "Pareto Network"
	case "POLY":
		return "Polymath"
	case "MTN":
		return "Medicalchain"
	case "JNT":
		return "Jibrel Network"
	case "CHSB":
		return "SwissBorg"
	case "ZLA":
		return "Zilla"
	case "ADB":
		return "adbank"
	case "HT":
		return "Huobi Token"
	case "DMT":
		return "DMarket"
	case "ING":
		return "Iungo"
	case "SWM":
		return "Swarm"
	case "TKY":
		return "THEKEY"
	case "DRPU":
		return "DCORP Utility"
	case "ESZ":
		return "EtherSportz"
	case "DXT":
		return "Datawallet"
	case "WPR":
		return "WePower"
	case "UCASH":
		return "UNIVERSAL CASH"
	case "MNTP":
		return "GoldMint"
	case "JEW":
		return "Shekel"
	case "MLM":
		return "MktCoin"
	case "AVH":
		return "Animation Vision Cash"
	case "LOCI":
		return "LOCIcoin"
	case "INDI":
		return "Indicoin"
	case "JC":
		return "Jesus Coin"
	case "BIO":
		return "BioCoin"
	case "SUP":
		return "Superior Coin"
	case "TIG":
		return "Tigereum"
	case "UTNP":
		return "Universa"
	case "ACAT":
		return "Alphacat"
	case "EVN":
		return "Envion"
	case "RMT":
		return "SureRemit"
	case "DTH":
		return "Dether"
	case "CAS":
		return "Cashaa"
	case "FSN":
		return "Fusion"
	case "W3C":
		return "W3Coin"
	case "ECH":
		return "Etherecash"
	case "MWAT":
		return "Restart Energy MWAT"
	case "GLA":
		return "Gladius Token"
	case "DADI":
		return "DADI"
	case "NTK":
		return "Neurotoken"
	case "GEM":
		return "Gems "
	case "NEC":
		return "Nectar"
	case "REN":
		return "Republic Protocol"
	case "LCC":
		return "Litecoin Cash"
	case "STQ":
		return "Storiqa"
	case "TDX":
		return "Tidex Token"
	case "CPY":
		return "COPYTRACK"
	case "NCASH":
		return "Nucleus Vision"
	case "ABT":
		return "Arcblock"
	case "REM":
		return "Remme"
	case "EXY":
		return "Experty"
	case "POA":
		return "POA Network"
	case "XNK":
		return "Ink Protocol"
	case "RKT":
		return "Rock"
	case "BEZ":
		return "Bezop"
	case "IHT":
		return "IHT Real Estate Protocol"
	case "RFR":
		return "Refereum"
	case "LYM":
		return "Lympo"
	case "SETH":
		return "Sether"
	case "CS":
		return "Credits"
	case "BEE":
		return "Bee Token"
	case "INSTAR":
		return "Insights Network"
	case "AUTO":
		return "Cube"
	case "EZT":
		return "EZToken"
	case "TUBE":
		return "BitTube"
	case "LEDU":
		return "Education Ecosystem"
	case "TUSD":
		return "TrueUSD"
	case "HQX":
		return "HOQU"
	case "STAC":
		return "StarterCoin"
	case "ONT":
		return "Ontology"
	case "DATX":
		return "DATx"
	case "J8T":
		return "JET8"
	case "CHP":
		return "CoinPoker"
	case "TOMO":
		return "TomoChain"
	case "GRFT":
		return "Graft"
	case "BAX":
		return "BABB"
	case "ELEC":
		return "Electrify.Asia"
	case "BTCP":
		return "Bitcoin Private"
	case "TEN":
		return "Tokenomy"
	case "RVN":
		return "Ravencoin"
	case "TFD":
		return "TE-FOOD"
	case "SHIP":
		return "ShipChain"
	case "LDC":
		return "Leadcoin"
	case "SHP":
		return "Sharpe Platform Token"
	case "LALA":
		return "LALA World"
	case "OCC":
		return "Octoin Coin"
	case "DEB":
		return "Debitum"
	case "CENNZ":
		return "Centrality"
	case "HAV":
		return "Havven"
	case "FLUZ":
		return "Fluz Fluz"
	case "LOOM":
		return "Loom Network"
	case "GETX":
		return "Guaranteed Ethurance Token Extra"
	case "HIRE":
		return "HireMatch"
	case "DROP":
		return "Dropil"
	case "BANCA":
		return "Banca"
	case "DRG":
		return "Dragon Coins"
	case "LATX":
		return "LatiumX"
	case "NANJ":
		return "NANJCOIN"
	case "CKUSD":
		return "CK USD"
	case "UP":
		return "UpToken"
	case "BBN":
		return "Banyan Network"
	case "NOAH":
		return "Noah Coin"
	case "LGO":
		return "LGO Exchange"
	case "1WO":
		return "1World"
	case "NPX":
		return "NaPoleonX"
	case "NPXS":
		return "Pundi X"
	case "BITG":
		return "Bitcoin Green"
	case "BFT":
		return "BnkToTheFuture"
	case "WAN":
		return "Wanchain"
	case "AMLT":
		return "AMLT"
	case "MITH":
		return "Mithril"
	case "LST":
		return "Lendroid Support Token"
	case "PCL":
		return "Peculium"
	case "SIG":
		return "Spectiv"
	case "RNTB":
		return "BitRent"
	case "XBP":
		return "BlitzPredict"
	case "LNC":
		return "Blocklancer"
	case "SPD":
		return "Stipend"
	case "IPSX":
		return "IP Exchange"
	case "SCC":
		return "StockChain"
	case "BSTN":
		return "BitStation"
	case "SWTH":
		return "Switcheo"
	case "SEN":
		return "Consensus"
	case "XCLR":
		return "ClearCoin"
	case "SENC":
		return "Sentinel Chain"
	case "VIT":
		return "Vice Industry Token"
	case "FDZ":
		return "Friendz"
	case "TPAY":
		return "TokenPay"
	case "BERRY":
		return "Rentberry"
	case "XTL":
		return "Stellite"
	case "NCT":
		return "PolySwarm"
	case "ODE":
		return "ODEM"
	case "XMO":
		return "Monero Original"
	case "XSN":
		return "Stakenet"
	case "XDCE":
		return "XinFin Network"
	case "TDS":
		return "TokenDesk"
	case "BBI":
		return "BelugaPay"
	case "FID":
		return "Fidelium"
	case "CTXC":
		return "Cortex"
	case "ATC":
		return "Arbitracoin"
	case "CPX":
		return "Apex"
	case "CVT":
		return "CyberVein"
	case "SENT":
		return "Sentinel"
	case "EOSDAC":
		return "eosDAC"
	case "UUU":
		return "U Network"
	case "ADH":
		return "AdHive"
	case "SNIP":
		return "SnipCoin"
	case "BSM":
		return "Bitsum"
	case "DEV":
		return "DeviantCoin"
	case "CBT":
		return "CommerceBlock"
	case "GRMD":
		return "GreenMed"
	case "AUC":
		return "Auctus"
	case "BUBO":
		return "Budbo"
	case "XMC":
		return "Monero Classic"
	case "DAN":
		return "Daneel"
	case "BRM":
		return "BrahmaOS"
	case "MFG":
		return "SyncFab"
	case "DIG":
		return "Dignity"
	case "ADI":
		return "Aditus"
	case "TRIO":
		return "Tripio"
	case "XHV":
		return "Haven Protocol"
	case "KST":
		return "StarCoin"
	case "DERO":
		return "Dero"
	case "EFX":
		return "Effect.AI"
	case "FTX":
		return "FintruX Network"
	case "EARTH":
		return "Earth Token"
	case "MRK":
		return "MARK.SPACE"
	case "CROP":
		return "Cropcoin"
	case "SRCOIN":
		return "SRCOIN"
	case "CHX":
		return "Own"
	case "MSR":
		return "Masari"
	case "DOCK":
		return "Dock"
	case "PHI":
		return "PHI Token"
	case "BBC":
		return "TraDove B2BCoin"
	case "DML":
		return "Decentralized Machine Learning"
	case "HBZ":
		return "Helbiz"
	case "ORI":
		return "Origami"
	case "TRAK":
		return "TrakInvest"
	case "APH":
		return "Aphelion"
	case "ZCO":
		return "Zebi"
	case "LND":
		return "Lendingblock"
	case "XES":
		return "Proxeus"
	case "VIPS":
		return "Vipstar Coin"
	case "RBLX":
		return "Rublix"
	case "BTRN":
		return "Biotron"
	case "PNT":
		return "Penta"
	case "NBAI":
		return "Nebula AI"
	case "LRN":
		return "Loopring NEO"
	case "NEXO":
		return "Nexo"
	case "VME":
		return "VeriME"
	case "DAX":
		return "DAEX"
	case "HYDRO":
		return "Hydro"
	case "SS":
		return "Sharder"
	case "CEL":
		return "Celsius"
	case "TTT":
		return "TrustNote"
	case "BCI":
		return "Bitcoin Interest"
	case "BETR":
		return "BetterBetting"
	case "TNS":
		return "Transcodium"
	case "AMN":
		return "Amon"
	case "PLAN":
		return "Plancoin"
	case "FLP":
		return "FLIP"
	case "CMCT":
		return "Crowd Machine"
	case "MITX":
		return "Morpheus Labs"
	case "LIVE":
		return "Live Stars"
	case "MTC":
		return "Docademic"
	case "MT":
		return "MyToken"
	case "NTY":
		return "Nexty"
	case "CJT":
		return "ConnectJob"
	case "GES":
		return "Galaxy eSolutions"
	case "BOUTS":
		return "BoutsPro"
	case "PAL":
		return "Pal Network"
	case "CRE":
		return "Cybereits"
	case "GENE":
		return "Parkgene"
	case "APR":
		return "APR Coin"
	case "AC3":
		return "AC3"
	case "FXT":
		return "FuzeX"
	case "ZIPT":
		return "Zippie"
	case "SKM":
		return "Skrumble Network"
	case "GEN":
		return "DAOstack"
	case "BZNT":
		return "Bezant"
	case "LIF":
		return "Winding Tree"
	case "TEAM":
		return "TEAM (TokenStars)"
	case "OOT":
		return "Utrum"
	case "FREC":
		return "Freyrchain"
	case "EDU":
		return "EduCoin"
	case "CNN":
		return "Content Neutrality Network"
	case "INSUR":
		return "InsurChain"
	case "GSC":
		return "Global Social Chain"
	case "SGCC":
		return "Super Game Chain"
	case "DGX":
		return "Digix Gold Token"
	case "INC":
		return "Influence Chain"
	case "IIC":
		return "Intelligent Investment Chain"
	case "SKB":
		return "Sakura Bloom"
	case "BANK":
		return "Bank Coin"
	case "NPER":
		return "NPER"
	case "JOINT":
		return "Joint Ventures"
	case "DASC":
		return "DasCoin"
	case "BMH":
		return "BlockMesh"
	case "LOKI":
		return "Loki"
	case "SGN":
		return "Signals Network"
	case "SHL":
		return "Oyster Shell"
	case "FND":
		return "FundRequest"
	case "DTRC":
		return "Datarius Credit"
	case "CLN":
		return "Colu Local Network"
	case "HER":
		return "HeroNode"
	case "TKA":
		return "Tokia"
	case "CLO":
		return "Callisto Network"
	case "UBT":
		return "Unibright"
	case "PAT":
		return "Patron"
	case "LBA":
		return "Cred"
	case "LWF":
		return "Local World Forwarders"
	case "OPEN":
		return "Open Platform"
	case "MRPH":
		return "Morpheus.Network"
	case "SNTR":
		return "Silent Notary"
	case "XYO":
		return "XYO Network"
	case "CPT":
		return "Cryptaur"
	case "APIS":
		return "APIS"
	case "FT":
		return "Fabric Token"
	case "XRH":
		return "Rhenium"
	case "CAZ":
		return "Cazcoin"
	case "DGTX":
		return "Digitex Futures"
	case "GIN":
		return "GINcoin"
	case "INV":
		return "Invacio"
	case "FACE":
		return "Faceter"
	case "AVA":
		return "Travala"
	case "IOTX":
		return "IoTeX"
	case "EXC":
		return "Eximchain"
	case "LUC":
		return "Level Up Coin"
	case "NKN":
		return "NKN"
	case "NAVI":
		return "Naviaddress"
	case "ZIP":
		return "Zipper"
	case "SOUL":
		return "Phantasma"
	case "REPO":
		return "REPO"
	case "SEELE":
		return "Seele"
	case "EJOY":
		return "EJOY"
	case "IVY":
		return "Ivy"
	case "CNET":
		return "ContractNet"
	case "BBO":
		return "Bigbom"
	case "0xBTC":
		return "0xBitcoin"
	case "PAI":
		return "PCHAIN"
	case "QKC":
		return "QuarkChain"
	case "LYL":
		return "LoyalCoin"
	case "BNK":
		return "Bankera"
	case "ETZ":
		return "Ether Zero"
	case "OMX":
		return "Shivom"
	case "MEDX":
		return "MediBloc ERC20"
	case "FTO":
		return "FuturoCoin"
	case "ABYSS":
		return "The Abyss"
	case "PMNT":
		return "Paymon"
	case "HUR":
		return "Hurify"
	case "TMT":
		return "TRAXIA"
	case "EGCC":
		return "Engine"
	case "MRQ":
		return "MIRQ"
	case "PKC":
		return "PikcioChain"
	case "CBC":
		return "CashBet Coin"
	case "CEEK":
		return "CEEK VR"
	case "SAL":
		return "SalPay"
	case "COU":
		return "Couchain"
	case "XMX":
		return "XMax"
	case "GO":
		return "GoChain"
	case "SSP":
		return "Smartshare"
	case "HOLD":
		return "HOLD"
	case "TRTT":
		return "Trittium"
	case "UPP":
		return "Sentinel Protocol"
	case "BWT":
		return "Bittwatt"
	case "DAG":
		return "Constellation"
	case "MVP":
		return "Merculet"
	case "FGC":
		return "FantasyGold"
	case "UCT":
		return "Ubique Chain Of Things"
	case "ETK":
		return "EnergiToken"
	case "MET":
		return "Metronome"
	case "AOA":
		return "Aurora"
	case "ALX":
		return "ALAX"
	case "TERN":
		return "Ternio"
	case "BOE":
		return "Bodhi ETH"
	case "ORS":
		return "Origin Sport"
	case "RTE":
		return "Rate3"
	case "DCC":
		return "Distributed Credit Chain"
	case "ZCN":
		return "0chain"
	case "ZINC":
		return "ZINC"
	case "FSBT":
		return "Forty Seven Bank"
	case "EGT":
		return "Egretia"
	case "TTU":
		return "TaTaTu"
	case "DOR":
		return "Dorado"
	case "CAR":
		return "CarBlock"
	case "BOB":
		return "Bob's Repair"
	case "KNDC":
		return "KanadeCoin"
	case "CARD":
		return "Cardstack"
	case "WWB":
		return "Wowbit"
	case "ONL":
		return "On.Live"
	case "OTB":
		return "OTCBTC Token"
	case "CONI":
		return "Coni"
	case "MFT":
		return "Mainframe"
	case "CCCX":
		return "Clipper Coin"
	case "GOT":
		return "GoNetwork"
	case "THRT":
		return "Thrive Token"
	case "FTI":
		return "FansTime"
	case "PCH":
		return "POPCHAIN"
	case "SEER":
		return "SEER"
	case "QURO":
		return "Qurito"
	case "ESS":
		return "Essentia"
	case "KBC":
		return "Karatgold Coin"
	case "HSC":
		return "HashCoin"
	case "LIKE":
		return "LikeCoin"
	case "YUP":
		return "Crowdholding"
	case "XSG":
		return "SnowGem"
	case "DTX":
		return "DaTa eXchange"
	case "BKBT":
		return "BeeKan"
	case "MOC":
		return "Moss Coin"
	case "NIM":
		return "Nimiq"
	case "SFU":
		return "Saifu"
	case "BZ":
		return "Bit-Z Token"
	case "DWS":
		return "DWS"
	case "ZXC":
		return "0xcert"
	case "OLT":
		return "OneLedger"
	case "ATMI":
		return "Atonomi"
	case "XMCT":
		return "XMCT"
	case "FNKOS":
		return "FNKOS"
	case "RCD":
		return "RECORD"
	case "PSM":
		return "PRASM"
	case "NUSD":
		return "nUSD"
	case "PLY":
		return "PlayCoin QRC20"
	case "TGAME":
		return "Truegame"
	case "IQ":
		return "Everipedia"
	case "ENGT":
		return "Engagement Token"
	case "NOBS":
		return "No BS Crypto"
	case "BMX":
		return "BitMart Token"
	case "KAN":
		return "BitKan"
	case "CDM":
		return "Condominium"
	case "VITE":
		return "VITE"
	case "GARD":
		return "Hashgard"
	case "SCRL":
		return "SCRL"
	case "SPX":
		return "Sp8de"
	case "CET":
		return "CoinEx Token"
	case "AEG":
		return "Aegeus"
	case "RPL":
		return "Rocket Pool"
	case "ELY":
		return "Elysian"
	case "BOX":
		return "ContentBox"
	case "PTT":
		return "Proton Token"
	case "SOP":
		return "SoPay"
	case "JOT":
		return "Jury.Online Token"
	case "KRL":
		return "Kryll"
	case "LEMO":
		return "LemoChain"
	case "GBC":
		return "Gold Bits Coin"
	case "BWX":
		return "Blue Whale Token"
	case "WYS":
		return "wys Token"
	case "COSM":
		return "Cosmo Coin"
	case "NRVE":
		return "Narrative"
	case "OLE":
		return "Olive"
	case "TRTL":
		return "Turtlecoin"
	case "WT":
		return "WeToken"
	case "TOTO":
		return "Tourist Token"
	case "RLX":
		return "Relex"
	case "CHEX":
		return "CHEX"
	case "VIEW":
		return "View"
	case "VCT":
		return "ValueCyberToken"
	case "VIKKY":
		return "VikkyToken"
	case "FOXT":
		return "Fox Trading"
	case "RTB":
		return "AB-Chain RTB"
	case "TOLL":
		return "Bridge Protocol"
	case "GVE":
		return "Globalvillage Ecosystem"
	case "LCS":
		return "LocalCoinSwap"
	case "ZPR":
		return "ZPER"
	case "EMPR":
		return "empowr coin"
	case "LPC":
		return "Lightpaycoin"
	case "FUNDZ":
		return "FundToken"
	case "RYO":
		return "Ryo Currency"
	case "ACED":
		return "AceD"
	case "LFT":
		return "Linfinity"
	case "WAB":
		return "WABnetwork"
	case "CSM":
		return "Consentium"
	case "MVL":
		return "Mass Vehicle Ledger"
	case "XXX":
		return "AdultChain"
	case "NCP":
		return "Newton Coin Project"
	case "DACC":
		return "DACC"
	case "TOS":
		return "ThingsOperatingSystem"
	case "PGN":
		return "Pigeoncoin"
	case "EURS":
		return "STASIS EURS"
	case "EXMR":
		return "Ethereum Monero"
	case "NIX":
		return "NIX"
	case "APL":
		return "Apollo Currency"
	case "HORUS":
		return "HorusPay"
	case "BIFI":
		return "Bitcoin File"
	case "BWS":
		return "Bitcoin W Spectrum"
	case "DPN":
		return "DIPNET"
	case "VEX":
		return "Vexanium"
	case "HDAC":
		return "Hdac"
	case "NPW":
		return "New Power Coin"
	case "KWH":
		return "KWHCoin"
	case "MCT":
		return "Master Contract Token"
	case "WSD":
		return "White Standard"
	case "ACDC":
		return "Volt"
	case "NBR":
		return "Niobio Cash"
	case "HRC":
		return "Haracoin"
	case "VIVID":
		return "Vivid Coin"
	case "PUREX":
		return "Pure"
	case "CEN":
		return "Coinsuper Ecosystem Network"
	case "BITX":
		return "BitScreener Token"
	case "VTHO":
		return "VeThor Token"
	case "PRIV":
		return "PRiVCY"
	case "RMESH":
		return "RightMesh"
	case "BBK":
		return "Brickblock"
	case "NCC":
		return "NeuroChain"
	case "COIN":
		return "Coinvest"
	case "KLKS":
		return "Kalkulus"
	case "LSTR":
		return "Luna Stars"
	case "BHPC":
		return "BHPCash"
	case "INCX":
		return "InternationalCryptoX"
	case "ZMN":
		return "ZMINE"
	case "SEM":
		return "Semux"
	case "ARO":
		return "Arionum"
	case "ACRE":
		return "ACRE"
	case "IOV":
		return "Carlive Chain"
	case "WEB":
		return "Webcoin"
	case "FMF":
		return "Formosa Financial"
	case "ZEL":
		return "ZelCash"
	case "BNN":
		return "BrokerNekoNetwork"
	case "OBT":
		return "Orbis Token"
	case "CZR":
		return "CanonChain"
	case "FOOD":
		return "FoodCoin"
	case "IDOL":
		return "IDOL COIN"
	case "PTS":
		return "PitisCoin"
	case "OPCX":
		return "OPCoinX"
	case "BND":
		return "Blocknode"
	case "XUN":
		return "UltraNote Coin"
	case "BTK":
		return "Bitcoin Token"
	case "ARLIZE":
		return "ARLIZE"
	case "DTEM":
		return "Dystem"
	case "ELI":
		return "Eligma Token"
	case "YOU":
		return "YOU COIN"
	case "DACS":
		return "DACSEE"
	case "EBC":
		return "EBCoin"
	case "TCH":
		return "Thore Cash"
	case "ADN":
		return "Adrenaline"
	case "SDA":
		return "Six Domain Chain"
	case "ESCO":
		return "EscrowCoin"
	case "YCC":
		return "Yuan Chain Coin"
	case "PC":
		return "Promotion Coin"
	case "GMCN":
		return "GambleCoin"
	case "VITAE":
		return "Vitae"
	case "EPLUS":
		return "EPLUS Coin"
	case "ROCK2":
		return "ICE ROCK MINING"
	case "BCV":
		return "BitCapitalVendor"
	case "XTRD":
		return "XTRD"
	case "BTCN":
		return "BitcoiNote"
	case "NAM":
		return "NAM COIN"
	case "LXT":
		return "Litex"
	case "EUNO":
		return "EUNO"
	case "MGD":
		return "MassGrid"
	case "EST":
		return "Esports Token"
	case "EXT":
		return "Experience Token"
	case "BTI":
		return "Bitcoin Instant"
	case "EDS":
		return "Endorsit"
	case "VET":
		return "Vechain"
	case "KIND":
		return "Kind Ads Token"
	case "X8X":
		return "X8X Token"
	case "CMM":
		return "Commercium"
	case "ECOM":
		return "Omnitude"
	case "VIN":
		return "VINchain"
	case "LINA":
		return "Lina"
	case "BTT":
		return "Blocktrade Token"
	case "INO":
		return "INO COIN"
	case "KNT":
		return "Kora Network Token"
	case "CROAT":
		return "CROAT"
	case "BTCONE":
		return "BitCoin One"
	case "AVINOC":
		return "AVINOC"
	case "WIKI":
		return "Wiki Token"
	case "SPN":
		return "Sapien"
	case "NUG":
		return "Nuggets"
	case "BBS":
		return "BBSCoin"
	case "SCR":
		return "Scorum Coins"
	case "NBC":
		return "Niobium Coin"
	case "NPXSXEM":
		return "Pundi X NEM"
	case "XOV":
		return "XOVBank"
	case "XCEL":
		return "XcelToken"
	case "LYNX":
		return "Lynx"
	case "UST":
		return "Ultra Salescloud"
	case "OPTI":
		return "OptiToken"
	case "GIC":
		return "Giant"
	case "ABDT":
		return "Atlantis Blue Digital Token"
	case "PKG":
		return "PKG Token"
	case "BOC":
		return "BingoCoin"
	case "HIGHT":
		return "HighCoin"
	case "RDC":
		return "Ordocoin"
	case "NEWOS":
		return "NewsToken"
	case "PDX":
		return "PayDay Coin"
	case "XPAT":
		return "Bitnation"
	case "ICR":
		return "InterCrone"
	case "RGS":
		return "RusGas"
	case "MXM":
		return "Maximine Coin"
	case "INB":
		return "Insight Chain"
	case "GIO":
		return "Graviocoin"
	case "SDS":
		return "Alchemint Standards"
	case "OWN":
		return "OWNDATA"
	case "IG":
		return "IGToken"
	case "HTH":
		return "Help The Homeless Coin"
	case "GSE":
		return "GSENetwork"
	case "DGS":
		return "Dragonglass"
	case "XDNA":
		return "XDNA"
	case "XPX":
		return "ProximaX"
	case "NYEX":
		return "Nyerium"
	case "FORK":
		return "Forkcoin"
	case "TIC":
		return "Thingschain"
	case "EGEM":
		return "EtherGem"
	case "AREPA":
		return "Arepacoin"
	case "XET":
		return "ETERNAL TOKEN"
	case "CEDEX":
		return "CEDEX Coin"
	case "MEETONE":
		return "MEET.ONE"
	case "KARMA":
		return "KARMA"
	case "NOKU":
		return "Noku"
	case "DX":
		return "DxChain Token"
	case "UBEX":
		return "Ubex"
	case "PASS":
		return "Blockpass"
	case "BAAS":
		return "BaaSid"
	case "PCO":
		return "Pecunio"
	case "THR":
		return "ThoreCoin"
	case "FRRN":
		return "Ferron"
	case "CYFM":
		return "CyberFM"
	case "HYC":
		return "HYCON"
	case "METM":
		return "MetaMorph"
	case "C2C":
		return "C2C System"
	case "AKA":
		return "Akroma"
	case "OBTC":
		return "Obitan Chain"
	case "TKT":
		return "Twinkle"
	case "DAC":
		return "Davinci Coin"
	case "QNT":
		return "Quant"
	case "ABL":
		return "Airbloc"
	case "SAC":
		return "Smart Application Chain"
	case "ZCR":
		return "ZCore"
	case "XAP":
		return "Apollon"
	case "IFP":
		return "Infinipay"
	case "SVD":
		return "savedroid"
	case "YLC":
		return "YoloCash"
	case "MERO":
		return "Mero"
	case "PMA":
		return "PumaPay"
	case "ARION":
		return "Arion"
	case "XBI":
		return "Bitcoin Incognito"
	case "SGP":
		return "SGPay"
	case "FTT":
		return "FarmaTrust"
	case "HYB":
		return "Hybrid Block"
	case "TALAO":
		return "Talao"
	case "HB":
		return "HeartBout"
	case "LGS":
		return "LogisCoin"
	case "TDC":
		return "Trendercoin"
	case "FNTB":
		return "Fintab"
	case "ALTX":
		return "Alttex"
	case "SEAL":
		return "Seal Network"
	case "LKY":
		return "Linkey"
	case "ABX":
		return "Arbidex"
	case "COMP":
		return "Compound Coin"
	case "HAND":
		return "ShowHand"
	case "HIT":
		return "HitChain"
	case "SC2":
		return "SecureCloudCoin"
	case "GPKR":
		return "Gold Poker"
	case "TWIST":
		return "TWIST"
	case "ZP":
		return "Zen Protocol"
	case "MOZO":
		return "Mozo Token"
	case "ECT":
		return "SuperEdge"
	case "MFTU":
		return "Mainstream For The Underground"
	case "DFS":
		return "Defense"
	case "RTL":
		return "Rentledger"
	case "CATO":
		return "CatoCoin"
	case "RRC":
		return "RRCoin"
	case "RATING":
		return "DPRating"
	case "CTC":
		return "Credit Tag Chain"
	case "KNOW":
		return "KNOW"
	case "GRPH":
		return "Graphcoin"
	case "KXC":
		return "KingXChain"
	case "NSD":
		return "Nasdacoin"
	case "PRTX":
		return "Printex"
	case "LOBS":
		return "Lobstex"
	case "RBMC":
		return "Rubex Money"
	case "VDG":
		return "VeriDocGlobal"
	case "YUKI":
		return "YUKI"
	case "KWATT":
		return "4NEW"
	case "MIB":
		return "MIB Coin"
	case "COTN":
		return "CottonCoin"
	case "BITF":
		return "BitF"
	case "SONIQ":
		return "Soniq"
	case "GTM":
		return "Gentarium"
	case "DELTA":
		return "DeltaChain"
	case "NRG":
		return "Energi"
	case "FTXT":
		return "FUTURAX"
	case "DAV":
		return "DAV Coin"
	case "ABLX":
		return "ABLE"
	case "BNC0":
		return "Bionic"
	case "DOW":
		return "DOWCOIN"
	case "QBIT":
		return "Qubitica"
	case "BTN":
		return "BitNewChain"
	case "DNZ":
		return "Adenz"
	case "TAC":
		return "Traceability Chain"
	case "CST":
		return "Cryptosolartech"
	case "VULC":
		return "VULCANO"
	case "BQT":
		return "Blockchain Quotations Index Token"
	case "STR":
		return "Staker"
	case "UT":
		return "Ulord"
	case "FLM":
		return "FolmCoin"
	case "WIN":
		return "WinToken"
	case "TMC":
		return "Timicoin"
	case "ESN":
		return "Ethersocial"
	case "FKX":
		return "Knoxstertoken"
	case "BEET":
		return "Beetle Coin"
	case "IMT":
		return "Moneytoken"
	case "MIC":
		return "Mindexcoin"
	case "UBC":
		return "Ubcoin Market"
	case "TSC":
		return "Thunderstake"
	case "FLOT":
		return "Fire Lotto"
	case "ALI":
		return "AiLink Token"
	case "USE":
		return "Usechain Token"
	case "ZBA":
		return "Zoomba"
	case "SHE":
		return "ShineChain"
	case "BLACK":
		return "eosBLACK"
	case "MRI":
		return "Mirai"
	case "CYMT":
		return "CyberMusic"
	case "BTR":
		return "Bitether"
	case "GZE":
		return "GazeCoin"
	case "BUT":
		return "BitUP Token"
	case "UC":
		return "YouLive Coin"
	case "AMO":
		return "AMO Coin"
	case "CCL":
		return "CYCLEAN"
	case "DIN":
		return "Dinero"
	case "DIT":
		return "Digital Insurance Token"
	case "HAVY":
		return "Havy"
	case "CARE":
		return "Carebit"
	case "PRJ":
		return "Project Coin"
	case "DART":
		return "DarexTravel"
	case "CIF":
		return "Crypto Improvement Fund"
	case "IMP":
		return "Ether Kingdoms Token"
	case "C2P":
		return "Coin2Play"
	case "C8":
		return "Carboneum C8 Token"
	case "ZEST":
		return "ZEST"
	case "SNO":
		return "SaveNode"
	case "VSC":
		return "vSportCoin"
	case "PENG":
		return "Penguin Coin"
	case "RTH":
		return "Rotharium"
	case "RET":
		return "RealTract"
	case "QNTU":
		return "Quanta Utility Token"
	case "TV":
		return "Ti-Value"
	case "FOIN":
		return "FOIN"
	case "GNR":
		return "Gainer"
	case "BIR":
		return "Birake"
	case "MEX":
		return "MEX"
	case "AAA":
		return "Abulaba"
	case "DAXT":
		return "Digital Asset Exchange Token"
	case "BEN":
		return "BitCoen"
	case "ELLI":
		return "Elliot Coin"
	case "BTXC":
		return "Bettex Coin"
	case "CIT":
		return "CariNet"
	case "OLMP":
		return "Olympic"
	case "BTAD":
		return "Bitcoin Adult"
	case "BU":
		return "BUMO"
	case "URALS":
		return "UralsCoin"
	case "XRT":
		return "XRT Token"
	case "BFF":
		return "BFFDoom"
	case "IHF":
		return "Invictus Hyperion Fund"
	case "UCN":
		return "UChain"
	case "CRBT":
		return "Cruisebit"
	case "MOLK":
		return "MobilinkToken"
	case "EDN":
		return "Eden"
	case "GUSD":
		return "Gemini Dollar"
	case "SPND":
		return "Spendcoin"
	case "XCG":
		return "Xchange"
	case "CCC":
		return "Concierge Coin"
	case "ALC":
		return "ALLCOIN"
	case "CSTL":
		return "Castle"
	case "EVI":
		return "Evimeria"
	case "CFL":
		return "CryptoFlow"
	case "BOXX":
		return "BOXX Token Blockparty"
	case "IOG":
		return "Playgroundz"
	case "AOG":
		return "smARTOFGIVING"
	case "CTRT":
		return "Cryptrust"
	case "TCN":
		return "TCOIN"
	case "BUNNY":
		return "BunnyToken"
	case "PYN":
		return "PAYCENT"
	case "PLURA":
		return "PluraCoin"
	case "ROX":
		return "Robotina"
	case "CHE":
		return "Crypto Harbor Exchange"
	case "SIX":
		return "SIX"
	case "CMIT":
		return "CMITCOIN"
	case "PAX":
		return "Paxos Standard Token"
	case "GOSS":
		return "Gossipcoin"
	case "SOL":
		return "Sola Token"
	case "XCASH":
		return "X-Cash"
	case "SHARD":
		return "Shard"
	case "IQN":
		return "IQeon"
	case "PAXEX":
		return "PAXEX"
	case "PGT":
		return "Puregold Token"
	case "MLC":
		return "Mallcoin"
	case "PHON":
		return "Phonecoin"
	case "ACTP":
		return "Archetypal Network"
	case "ANON":
		return "ANON"
	case "ECOREAL":
		return "Ecoreal Estate"
	case "DAPS":
		return "DAPS Token"
	case "VSTR":
		return "Vestoria"
	case "CARAT":
		return "CARAT"
	case "MNP":
		return "MNPCoin"
	case "PYX":
		return "PyrexCoin"
	case "DACH":
		return "DACH Coin"
	case "ZB":
		return "ZB"
	case "MAS":
		return "MidasProtocol"
	case "WELL":
		return "WELL"
	case "TRXC":
		return "TRONCLASSIC"
	case "AZART":
		return "Azart"
	case "TMTG":
		return "The Midas Touch Gold"
	case "DAGT":
		return "Digital Asset Guarantee Token"
	case "HSN":
		return "Helper Search Token"
	case "WIT":
		return "WITChain"
	case "ERT":
		return "Eristica"
	case "AUX":
		return "Auxilium"
	case "WXC":
		return "WXCOINS"
	case "CPLO":
		return "Cpollo"
	case "SINS":
		return "SafeInsure"
	case "CRD":
		return "CryptalDash"
	case "NER":
		return "Nerves"
	case "MIR":
		return "MIR COIN"
	case "RPM":
		return "Repme"
	case "BETHER":
		return "Bethereum"
	case "RAGNA":
		return "Ragnarok"
	case "DEC":
		return "Darico Ecosystem Coin"
	case "XGS":
		return "GenesisX"
	case "LABH":
		return "Labh Coin"
	case "WBL":
		return "WIZBL"
	case "CIV":
		return "Civitas"
	case "BCARD":
		return "CARDbuyers"
	case "BENZ":
		return "Benz"
	case "XG":
		return "GIGA"
	case "ACM":
		return "Actinium"
	case "BLAST":
		return "BLAST"
	case "FREE":
		return "FREE Coin"
	case "TOL":
		return "Tolar"
	case "QUAN":
		return "Quantis Network"
	case "FOX":
		return "SmartFox"
	case "MASH":
		return "MASTERNET"
	case "OUR":
		return "Ourcoin"
	case "STEEP":
		return "SteepCoin"
	case "NRP":
		return "Neural Protocol"
	case "SCRIV":
		return "SCRIV NETWORK"
	case "WSP":
		return "Wispr"
	case "X12":
		return "X12 Coin"
	case "SHADE":
		return "SHADE Token"
	case "IFOOD":
		return "Ifoods Chain"
	case "EGX":
		return "EagleX"
	case "WIX":
		return "Wixlar"
	case "PNDM":
		return "Pandemia"
	case "BC":
		return "Block-Chain.com"
	case "RSTR":
		return "Ondori"
	case "USDC":
		return "USD//Coin"
	case "ETA":
		return "Etheera"
	case "BSX":
		return "Bitspace"
	case "WTL":
		return "Welltrado"
	case "SIM":
		return "Simmitri"
	case "NDX":
		return "nDEX"
	case "ZEUS":
		return "ZeusCrowdfunding"
	case "BCZERO":
		return "Buggyra Coin Zero"
	case "WAGE":
		return "Digiwage"
	case "F1C":
		return "Future1coin"
	case "META":
		return "Metadium"
	case "QAC":
		return "Quasarcoin"
	case "COBRA":
		return "Cobrabytes"
	case "PSC":
		return "PrimeStone"
	case "SHPING":
		return "SHPING"
	case "S":
		return "Sharpay"
	case "QNO":
		return "QYNO"
	case "AEC":
		return "EmaratCoin"
	case "INCO":
		return "Incodium"
	case "AGLT":
		return "Agrolot"
	case "JIYOX":
		return "JIYO"
	case "SURE":
		return "SURETY"
	case "ICNQ":
		return "Iconiq Lab Token"
	case "RPD":
		return "Rapids"
	case "ENTS":
		return "EUNOMIA"
	case "SOOM":
		return "SOOM"
	case "SNET":
		return "Snetwork"
	case "SMQ":
		return "SIMDAQ"
	case "ABBC":
		return "Alibabacoin"
	case "OXY":
		return "Oxycoin"
	case "DEAL":
		return "iDealCash"
	case "WIRE":
		return "AirWire"
	case "DIVI":
		return "Divi"
	case "XIND":
		return "INDINODE"
	case "HUZU":
		return "HUZU"
	case "KUN":
		return "KUN"
	case "ZNT":
		return "Zenswap Network Token"
	case "ATH":
		return "Atheios"
	case "CDC":
		return "Commerce Data Connection"
	case "MMO":
		return "MMOCoin"
	case "SZC":
		return "ShopZcoin"
	case "BLOC":
		return "BLOC.MONEY"
	case "ETHO":
		return "Ether-1"
	case "CJS":
		return "CJs"
	case "DATP":
		return "Decentralized Asset Trading Platform"
	case "DEEX":
		return "DEEX"
	case "PLUS1":
		return "PlusOneCoin"
	case "IRD":
		return "Iridium"
	case "ZT":
		return "ZTCoin"
	case "HELP":
		return "GoHelpFund"
	case "IXE":
		return "IXTUS Edutainment"
	case "RPI":
		return "RPICoin"
	case "CHEESE":
		return "Cheesecoin"
	case "SOFR":
		return "Secured Overnight Financing Rate"
	default:
		return NameForSymbolManuallyAdded(symbol)
	}
}
