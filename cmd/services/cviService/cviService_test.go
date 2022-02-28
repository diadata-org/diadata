package main

import (
	filters "github.com/diadata-org/diadata/internal/pkg/filtersOptionService"
	"testing"
)

func TestCalculateMidAndDifference(t *testing.T) {
	near := CalculateMidAndDifference(getNearTermTestData())
	if 2.099999999999998 != near[1965].Difference {
		t.Error("Diff is", near[1965].Difference)
	}

	if 23.200000000000003 != near[1940].Difference {
		t.Error("Diff is", near[1940].Difference)
	}
}

//strike price at which the absolute difference between the call and put
//prices is smallest
func TestMinimim(t *testing.T) {
	near := CalculateMidAndDifference(getNearTermTestData())
	minimum := FindMinimumMid(near)
	if 1965 != minimum {
		t.Error(" minimum", minimum)
	}
}

func TestForward(t *testing.T) {
	near := CalculateMidAndDifference(getNearTermTestData())
	minimum := FindMinimumMid(near)
	forward := filters.CalculateForwardIndex(near[minimum].StrikePrice, 0.000305,  0.0683486, near[minimum].CallMid, near[minimum].PutMid)
	if 1962.8999562222655 != forward {
		t.Error("Forward ", forward)
	}
}

func TestFindk1(t *testing.T){
	var nearstrikePrices []float64

	near := CalculateMidAndDifference(getNearTermTestData())
	minimum := FindMinimumMid(near)
	forward := filters.CalculateForwardIndex(near[minimum].StrikePrice, 0.000305,  0.0683486, near[minimum].CallMid, near[minimum].PutMid)
	 for sp,_ := range near{
		 nearstrikePrices = append(nearstrikePrices,sp)
	 }

	 k1 := findK1(nearstrikePrices,forward)
	if 1960 != k1 {
		t.Error("k1 ", k1)
	}

}

func getNearTermTestData() map[float64]OptionsTable {
	var testData = make(map[float64]OptionsTable)
	testData[800] = OptionsTable{CallBid: 1160.90, CallAsk: 1164.40, PutBid: 0.00, PutAsk: 0.10}
	testData[900] = OptionsTable{CallBid: 1160.90, CallAsk: 1064.50, PutBid: 0.00, PutAsk: 0.10}
	testData[1000] = OptionsTable{CallBid: 961.00, CallAsk: 964.50, PutBid: 0.00, PutAsk: 0.10}
	testData[1050] = OptionsTable{CallBid: 911.00, CallAsk: 914.50, PutBid: 0.00, PutAsk: 0.10}
	testData[1100] = OptionsTable{CallBid: 861.00, CallAsk: 864.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1125] = OptionsTable{CallBid: 836.00, CallAsk: 839.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1150] = OptionsTable{CallBid: 811.00, CallAsk: 814.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1175] = OptionsTable{CallBid: 786.10, CallAsk: 789.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1200] = OptionsTable{CallBid: 761.10, CallAsk: 764.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1220] = OptionsTable{CallBid: 741.10, CallAsk: 744.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1225] = OptionsTable{CallBid: 736.10, CallAsk: 739.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1240] = OptionsTable{CallBid: 721.10, CallAsk: 724.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1250] = OptionsTable{CallBid: 711.10, CallAsk: 714.60, PutBid: 0.00, PutAsk: 0.05}
	testData[1260] = OptionsTable{CallBid: 701.10, CallAsk: 704.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1270] = OptionsTable{CallBid: 691.10, CallAsk: 694.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1275] = OptionsTable{CallBid: 686.10, CallAsk: 689.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1280] = OptionsTable{CallBid: 681.10, CallAsk: 684.60, PutBid: 0.00, PutAsk: 0.10}
	testData[1290] = OptionsTable{CallBid: 671.10, CallAsk: 674.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1300] = OptionsTable{CallBid: 661.10, CallAsk: 664.70, PutBid: 0.05, PutAsk: 0.10}
	testData[1305] = OptionsTable{CallBid: 656.10, CallAsk: 659.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1310] = OptionsTable{CallBid: 651.10, CallAsk: 654.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1315] = OptionsTable{CallBid: 646.10, CallAsk: 649.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1320] = OptionsTable{CallBid: 641.20, CallAsk: 644.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1325] = OptionsTable{CallBid: 636.20, CallAsk: 639.70, PutBid: 0.05, PutAsk: 0.10}
	testData[1330] = OptionsTable{CallBid: 631.20, CallAsk: 634.70, PutBid: 0.00, PutAsk: 0.10}
	testData[1335] = OptionsTable{CallBid: 626.20, CallAsk: 629.70, PutBid: 0.00, PutAsk: 0.15}
	testData[1340] = OptionsTable{CallBid: 621.20, CallAsk: 624.70, PutBid: 0.00, PutAsk: 0.15}
	testData[1345] = OptionsTable{CallBid: 616.20, CallAsk: 619.70, PutBid: 0.00, PutAsk: 0.15}
	testData[1350] = OptionsTable{CallBid: 611.20, CallAsk: 614.70, PutBid: 0.05, PutAsk: 0.15}
	testData[1355] = OptionsTable{CallBid: 606.20, CallAsk: 609.70, PutBid: 0.05, PutAsk: 0.35}
	testData[1360] = OptionsTable{CallBid: 601.20, CallAsk: 604.70, PutBid: 0.00, PutAsk: 0.35}
	testData[1365] = OptionsTable{CallBid: 596.20, CallAsk: 599.70, PutBid: 0.00, PutAsk: 0.35}
	testData[1370] = OptionsTable{CallBid: 591.20, CallAsk: 594.70, PutBid: 0.05, PutAsk: 0.35}
	testData[1375] = OptionsTable{CallBid: 586.20, CallAsk: 589.70, PutBid: 0.10, PutAsk: 0.15}
	testData[1380] = OptionsTable{CallBid: 581.20, CallAsk: 584.70, PutBid: 0.10, PutAsk: 0.20}
	testData[1385] = OptionsTable{CallBid: 576.20, CallAsk: 579.70, PutBid: 0.10, PutAsk: 0.35}
	testData[1390] = OptionsTable{CallBid: 571.20, CallAsk: 574.70, PutBid: 0.10, PutAsk: 0.35}
	testData[1395] = OptionsTable{CallBid: 566.20, CallAsk: 569.70, PutBid: 0.10, PutAsk: 0.15}
	testData[1400] = OptionsTable{CallBid: 561.20, CallAsk: 564.80, PutBid: 0.10, PutAsk: 0.15}
	testData[1405] = OptionsTable{CallBid: 556.20, CallAsk: 559.80, PutBid: 0.00, PutAsk: 0.35}
	testData[1410] = OptionsTable{CallBid: 551.20, CallAsk: 554.80, PutBid: 0.05, PutAsk: 0.40}
	testData[1415] = OptionsTable{CallBid: 546.20, CallAsk: 549.80, PutBid: 0.00, PutAsk: 0.40}
	testData[1420] = OptionsTable{CallBid: 541.20, CallAsk: 544.80, PutBid: 0.05, PutAsk: 0.40}
	testData[1425] = OptionsTable{CallBid: 536.30, CallAsk: 539.80, PutBid: 0.15, PutAsk: 0.20}
	testData[1430] = OptionsTable{CallBid: 531.30, CallAsk: 534.80, PutBid: 0.05, PutAsk: 0.40}
	testData[1435] = OptionsTable{CallBid: 526.30, CallAsk: 529.80, PutBid: 0.15, PutAsk: 0.40}
	testData[1440] = OptionsTable{CallBid: 521.30, CallAsk: 524.80, PutBid: 0.05, PutAsk: 0.30}
	testData[1445] = OptionsTable{CallBid: 516.30, CallAsk: 519.80, PutBid: 0.05, PutAsk: 0.40}
	testData[1450] = OptionsTable{CallBid: 511.30, CallAsk: 514.80, PutBid: 0.15, PutAsk: 0.25}
	testData[1455] = OptionsTable{CallBid: 506.30, CallAsk: 509.80, PutBid: 0.05, PutAsk: 0.45}
	testData[1460] = OptionsTable{CallBid: 501.30, CallAsk: 504.80, PutBid: 0.05, PutAsk: 0.45}
	testData[1465] = OptionsTable{CallBid: 496.30, CallAsk: 499.80, PutBid: 0.05, PutAsk: 0.45}
	testData[1470] = OptionsTable{CallBid: 491.30, CallAsk: 494.80, PutBid: 0.05, PutAsk: 0.45}
	testData[1475] = OptionsTable{CallBid: 486.30, CallAsk: 489.90, PutBid: 0.15, PutAsk: 0.25}
	testData[1480] = OptionsTable{CallBid: 481.30, CallAsk: 484.90, PutBid: 0.05, PutAsk: 0.45}
	testData[1485] = OptionsTable{CallBid: 476.30, CallAsk: 479.90, PutBid: 0.20, PutAsk: 0.50}
	testData[1490] = OptionsTable{CallBid: 471.30, CallAsk: 474.90, PutBid: 0.05, PutAsk: 0.30}
	testData[1495] = OptionsTable{CallBid: 466.40, CallAsk: 469.90, PutBid: 0.05, PutAsk: 0.50}
	testData[1500] = OptionsTable{CallBid: 461.40, CallAsk: 464.90, PutBid: 0.25, PutAsk: 0.40}
	testData[1505] = OptionsTable{CallBid: 456.40, CallAsk: 459.90, PutBid: 0.30, PutAsk: 0.35}
	testData[1510] = OptionsTable{CallBid: 451.40, CallAsk: 454.90, PutBid: 0.05, PutAsk: 0.55}
	testData[1515] = OptionsTable{CallBid: 446.40, CallAsk: 449.90, PutBid: 0.05, PutAsk: 0.55}
	testData[1520] = OptionsTable{CallBid: 441.40, CallAsk: 445.00, PutBid: 0.10, PutAsk: 0.60}
	testData[1525] = OptionsTable{CallBid: 436.40, CallAsk: 440.00, PutBid: 0.30, PutAsk: 0.40}
	testData[1530] = OptionsTable{CallBid: 431.40, CallAsk: 435.00, PutBid: 0.05, PutAsk: 0.60}
	testData[1535] = OptionsTable{CallBid: 426.40, CallAsk: 430.00, PutBid: 0.10, PutAsk: 0.65}
	testData[1540] = OptionsTable{CallBid: 421.40, CallAsk: 425.00, PutBid: 0.10, PutAsk: 0.65}
	testData[1545] = OptionsTable{CallBid: 416.50, CallAsk: 420.00, PutBid: 0.10, PutAsk: 0.65}
	testData[1550] = OptionsTable{CallBid: 411.50, CallAsk: 415.00, PutBid: 0.30, PutAsk: 0.70}
	testData[1555] = OptionsTable{CallBid: 406.50, CallAsk: 410.10, PutBid: 0.15, PutAsk: 0.70}
	testData[1560] = OptionsTable{CallBid: 401.50, CallAsk: 405.10, PutBid: 0.15, PutAsk: 0.70}
	testData[1565] = OptionsTable{CallBid: 396.50, CallAsk: 400.10, PutBid: 0.15, PutAsk: 0.70}
	testData[1570] = OptionsTable{CallBid: 391.50, CallAsk: 395.10, PutBid: 0.20, PutAsk: 0.75}
	testData[1575] = OptionsTable{CallBid: 386.50, CallAsk: 390.10, PutBid: 0.35, PutAsk: 0.75}
	testData[1580] = OptionsTable{CallBid: 381.50, CallAsk: 385.10, PutBid: 0.25, PutAsk: 0.80}
	testData[1585] = OptionsTable{CallBid: 376.60, CallAsk: 380.20, PutBid: 0.25, PutAsk: 0.80}
	testData[1590] = OptionsTable{CallBid: 371.60, CallAsk: 375.20, PutBid: 0.25, PutAsk: 0.80}
	testData[1595] = OptionsTable{CallBid: 366.60, CallAsk: 370.20, PutBid: 0.25, PutAsk: 0.80}
	testData[1600] = OptionsTable{CallBid: 361.60, CallAsk: 365.20, PutBid: 0.50, PutAsk: 0.85}
	testData[1605] = OptionsTable{CallBid: 356.60, CallAsk: 360.30, PutBid: 0.30, PutAsk: 0.85}
	testData[1610] = OptionsTable{CallBid: 351.60, CallAsk: 355.30, PutBid: 0.35, PutAsk: 0.90}
	testData[1620] = OptionsTable{CallBid: 341.70, CallAsk: 345.30, PutBid: 0.35, PutAsk: 0.90}
	testData[1625] = OptionsTable{CallBid: 336.70, CallAsk: 340.40, PutBid: 0.40, PutAsk: 0.95}
	testData[1630] = OptionsTable{CallBid: 331.70, CallAsk: 335.40, PutBid: 0.40, PutAsk: 0.95}
	testData[1635] = OptionsTable{CallBid: 326.70, CallAsk: 330.40, PutBid: 0.45, PutAsk: 0.95}
	testData[1635] = OptionsTable{CallBid: 326.70, CallAsk: 330.40, PutBid: 0.45, PutAsk: 1.00}
	testData[1640] = OptionsTable{CallBid: 321.80, CallAsk: 325.40, PutBid: 0.45, PutAsk: 1.00}
	testData[1645] = OptionsTable{CallBid: 316.80, CallAsk: 320.50, PutBid: 0.50, PutAsk: 1.05}
	testData[1650] = OptionsTable{CallBid: 311.80, CallAsk: 315.50, PutBid: 0.50, PutAsk: 0.85}
	testData[1655] = OptionsTable{CallBid: 306.80, CallAsk: 310.50, PutBid: 0.55, PutAsk: 1.10}
	testData[1660] = OptionsTable{CallBid: 301.90, CallAsk: 305.60, PutBid: 0.55, PutAsk: 1.10}
	testData[1665] = OptionsTable{CallBid: 296.90, CallAsk: 300.60, PutBid: 0.60, PutAsk: 1.15}
	testData[1670] = OptionsTable{CallBid: 291.90, CallAsk: 295.70, PutBid: 0.60, PutAsk: 1.15}
	testData[1675] = OptionsTable{CallBid: 287.00, CallAsk: 290.70, PutBid: 0.65, PutAsk: 1.20}
	testData[1680] = OptionsTable{CallBid: 282.00, CallAsk: 285.70, PutBid: 0.70, PutAsk: 1.25}
	testData[1685] = OptionsTable{CallBid: 277.00, CallAsk: 280.80, PutBid: 0.75, PutAsk: 1.30}
	testData[1690] = OptionsTable{CallBid: 272.10, CallAsk: 275.80, PutBid: 0.75, PutAsk: 1.30}
	testData[1695] = OptionsTable{CallBid: 267.10, CallAsk: 270.90, PutBid: 0.80, PutAsk: 1.35}
	testData[1700] = OptionsTable{CallBid: 262.10, CallAsk: 265.90, PutBid: 0.85, PutAsk: 1.40}
	testData[1705] = OptionsTable{CallBid: 257.20, CallAsk: 261.00, PutBid: 0.85, PutAsk: 1.40}
	testData[1710] = OptionsTable{CallBid: 252.20, CallAsk: 256.00, PutBid: 0.90, PutAsk: 1.45}
	testData[1715] = OptionsTable{CallBid: 247.30, CallAsk: 251.10, PutBid: 0.95, PutAsk: 1.50}
	testData[1720] = OptionsTable{CallBid: 242.30, CallAsk: 246.10, PutBid: 1.00, PutAsk: 1.55}
	testData[1725] = OptionsTable{CallBid: 237.40, CallAsk: 241.20, PutBid: 1.05, PutAsk: 1.60}
	testData[1730] = OptionsTable{CallBid: 232.40, CallAsk: 236.30, PutBid: 1.10, PutAsk: 1.65}
	testData[1735] = OptionsTable{CallBid: 227.50, CallAsk: 231.30, PutBid: 1.15, PutAsk: 1.70}
	testData[1740] = OptionsTable{CallBid: 222.50, CallAsk: 226.40, PutBid: 1.20, PutAsk: 1.75}
	testData[1745] = OptionsTable{CallBid: 217.60, CallAsk: 221.50, PutBid: 1.25, PutAsk: 1.85}
	testData[1750] = OptionsTable{CallBid: 212.60, CallAsk: 216.60, PutBid: 1.30, PutAsk: 1.90}
	testData[1755] = OptionsTable{CallBid: 207.70, CallAsk: 211.60, PutBid: 1.40, PutAsk: 1.95}
	testData[1760] = OptionsTable{CallBid: 202.80, CallAsk: 206.70, PutBid: 1.45, PutAsk: 2.05}
	testData[1765] = OptionsTable{CallBid: 197.80, CallAsk: 201.80, PutBid: 1.50, PutAsk: 2.15}
	testData[1770] = OptionsTable{CallBid: 192.90, CallAsk: 196.90, PutBid: 1.60, PutAsk: 2.20}
	testData[1775] = OptionsTable{CallBid: 188.00, CallAsk: 192.00, PutBid: 1.65, PutAsk: 2.35}
	testData[1780] = OptionsTable{CallBid: 183.10, CallAsk: 187.10, PutBid: 1.75, PutAsk: 2.40}
	testData[1785] = OptionsTable{CallBid: 178.20, CallAsk: 182.20, PutBid: 1.85, PutAsk: 2.50}
	testData[1790] = OptionsTable{CallBid: 173.30, CallAsk: 177.30, PutBid: 1.90, PutAsk: 2.60}
	testData[1795] = OptionsTable{CallBid: 168.40, CallAsk: 172.40, PutBid: 2.00, PutAsk: 2.75}
	testData[1800] = OptionsTable{CallBid: 163.50, CallAsk: 167.50, PutBid: 2.15, PutAsk: 2.90}
	testData[1805] = OptionsTable{CallBid: 158.60, CallAsk: 162.60, PutBid: 2.25, PutAsk: 3.00}
	testData[1810] = OptionsTable{CallBid: 153.80, CallAsk: 157.80, PutBid: 2.35, PutAsk: 3.20}
	testData[1815] = OptionsTable{CallBid: 148.90, CallAsk: 152.90, PutBid: 2.50, PutAsk: 3.40}
	testData[1820] = OptionsTable{CallBid: 144.10, CallAsk: 148.10, PutBid: 2.65, PutAsk: 3.50}
	testData[1825] = OptionsTable{CallBid: 139.20, CallAsk: 143.30, PutBid: 3.00, PutAsk: 3.60}
	testData[1825] = OptionsTable{CallBid: 139.20, CallAsk: 143.30, PutBid: 3.00, PutAsk: 3.60}
	testData[1830] = OptionsTable{CallBid: 134.40, CallAsk: 138.40, PutBid: 3.00, PutAsk: 3.90}
	testData[1835] = OptionsTable{CallBid: 129.60, CallAsk: 133.60, PutBid: 3.20, PutAsk: 4.10}
	testData[1840] = OptionsTable{CallBid: 124.80, CallAsk: 128.80, PutBid: 3.40, PutAsk: 4.40}
	testData[1845] = OptionsTable{CallBid: 120.10, CallAsk: 124.10, PutBid: 3.60, PutAsk: 4.60}
	testData[1850] = OptionsTable{CallBid: 115.40, CallAsk: 119.30, PutBid: 3.80, PutAsk: 4.90}
	testData[1855] = OptionsTable{CallBid: 110.60, CallAsk: 114.60, PutBid: 4.10, PutAsk: 5.20}
	testData[1860] = OptionsTable{CallBid: 105.90, CallAsk: 109.90, PutBid: 4.40, PutAsk: 5.50}
	testData[1865] = OptionsTable{CallBid: 101.30, CallAsk: 105.20, PutBid: 4.70, PutAsk: 5.80}
	testData[1870] = OptionsTable{CallBid: 96.60, CallAsk: 100.50, PutBid: 5.00, PutAsk: 6.20}
	testData[1875] = OptionsTable{CallBid: 92.00, CallAsk: 95.90, PutBid: 5.40, PutAsk: 6.60}
	testData[1880] = OptionsTable{CallBid: 87.40, CallAsk: 91.30, PutBid: 5.80, PutAsk: 7.00}
	testData[1885] = OptionsTable{CallBid: 82.90, CallAsk: 86.70, PutBid: 6.20, PutAsk: 7.50}
	testData[1890] = OptionsTable{CallBid: 78.40, CallAsk: 82.20, PutBid: 6.70, PutAsk: 8.00}
	testData[1895] = OptionsTable{CallBid: 74.00, CallAsk: 77.70, PutBid: 7.20, PutAsk: 8.60}
	testData[1900] = OptionsTable{CallBid: 69.60, CallAsk: 73.20, PutBid: 7.80, PutAsk: 8.80}
	testData[1905] = OptionsTable{CallBid: 66.00, CallAsk: 68.50, PutBid: 8.50, PutAsk: 9.50}
	testData[1910] = OptionsTable{CallBid: 61.60, CallAsk: 64.10, PutBid: 9.10, PutAsk: 10.20}
	testData[1915] = OptionsTable{CallBid: 57.40, CallAsk: 59.80, PutBid: 9.90, PutAsk: 11.30}
	testData[1920] = OptionsTable{CallBid: 53.30, CallAsk: 55.60, PutBid: 10.70, PutAsk: 12.10}
	testData[1925] = OptionsTable{CallBid: 49.10, CallAsk: 51.20, PutBid: 11.60, PutAsk: 12.60}
	testData[1930] = OptionsTable{CallBid: 45.20, CallAsk: 47.30, PutBid: 12.50, PutAsk: 14.00}
	testData[1935] = OptionsTable{CallBid: 41.20, CallAsk: 43.40, PutBid: 13.60, PutAsk: 14.70}
	testData[1940] = OptionsTable{CallBid: 37.40, CallAsk: 39.50, PutBid: 14.70, PutAsk: 15.80}
	testData[1945] = OptionsTable{CallBid: 33.70, CallAsk: 35.70, PutBid: 15.90, PutAsk: 17.20}
	testData[1950] = OptionsTable{CallBid: 30.10, CallAsk: 32.10, PutBid: 17.70, PutAsk: 18.80}
	testData[1955] = OptionsTable{CallBid: 26.70, CallAsk: 28.50, PutBid: 19.00, PutAsk: 20.50}
	testData[1960] = OptionsTable{CallBid: 23.40, CallAsk: 25.10, PutBid: 20.60, PutAsk: 22.00}
	testData[1965] = OptionsTable{CallBid: 20.30, CallAsk: 21.80, PutBid: 22.30, PutAsk: 24.00}
	testData[1970] = OptionsTable{CallBid: 17.40, CallAsk: 18.80, PutBid: 24.30, PutAsk: 25.80}
	testData[1975] = OptionsTable{CallBid: 14.60, CallAsk: 15.90, PutBid: 26.50, PutAsk: 28.10}

	//1975 14.60 15.90 26.50 28.10
	//1980 12.20 13.30 28.90 30.60
	//1985 9.90 11.00 31.40 33.20
	//1990 7.90 9.00 34.30 36.50
	//1995 6.20 7.10 37.40 39.70
	//2000 4.70 5.20 40.70 43.20
	//2005 3.40 4.20 44.00 47.70
	//2010 2.65 3.10 48.00 51.40
	//2015 1.75 2.30 52.20 56.00
	//2020 1.20 1.70 56.60 60.40
	//2025 1.00 1.25 61.20 65.00
	//2030 0.45 1.00 65.90 69.70
	//2035 0.25 0.80 70.70 74.40
	//2040 0.35 0.65 75.60 79.30
	//2045 0.20 0.60 80.50 84.10
	//2050 0.20 0.30 85.40 89.00
	//2055 0.15 0.50 90.40 94.00
	//2060 0.15 0.30 95.30 98.90
	//2065 0.15 0.20 100.30 103.90
	//2070 0.10 0.20 105.30 108.90
	//2075 0.10 0.20 110.30 113.80
	//2080 0.05 0.45 115.30 118.80
	//2085 0.05 0.40 120.30 123.80
	//2090 0.05 0.15 125.30 128.80
	//2095 0.05 0.35 130.30 133.80
	//2100 0.05 0.15 135.30 138.80
	//2120 0.00 0.15 155.30 158.80
	//2125 0.05 0.15 160.30 163.80
	//2150 0.00 0.10 185.20 188.80
	//2175 0.00 0.05 210.20 213.70
	//2200 0.00 0.05 235.20 238.70
	//2225 0.05 0.10 260.20 263.70
	//2250 0.00 0.05 285.20 288.70

	return testData

}
