# Crypto Volatility Index

## Methodology

The CVI is calculated on top of bid/ask information received from crypto \(BTC\) option markets. First of all, the current level of the orderbook is recorded and placed into a timestamped order. Instruments for each strike price are monitored, both for Call and Put options.

Instruments are named uniquely to identify them. For example, BTC-03JUN2020-6000-C is the name of a Bitcoin call option expiring at 3rd of June 2020 for a strike price of 6000 USD.

To calculate the CVI, the option levels of two dates are considered: The next Friday that is at least one month away from now and the first Friday with available options before that date. With that we ensure that the future market is observed appropriately.

{% hint style="info" %}
Options typically expire on Fridays only, thus we can only observe them for these dates.
{% endhint %}

The CVI is then calculated using this formula:

$$
x = 100 \cdot \sqrt{\{T_1\sigma_1^2[\frac{N_{T_2} - N_{30}}{N_{T_2}-N_{T_1}}]+T_2\sigma^2_2[\frac{N_{30} - N_{T_1}}{N_{T_2}-N_{T_1}}]\}\cdot\frac{N_{365}}{N_{30}}}
$$

where N describe the minutes until the respective event \(e.g. minutes in a year, 30 days or minutes until expiry of option T1\). T1 and T2 are the fractions of minutes left until expiry of these options compared to the minutes in a full year.

The variances are calculated as such:

$$
\sigma^2=\frac{2}{T}\sum_{i}\frac{\Delta K_i}{K^2_i}e^{RT}Q(K_i)
$$

with the sum of deltas being defined as the intervals between strike prices for each strike price. R is the risk-free lending rate and Q\(k\) is used to describe the average of the bid-ask-spread for an option K\_i.

The CVI starts at a value of 1000.

