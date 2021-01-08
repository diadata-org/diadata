---
description: A basic introduction to some fundamental concepts in crypto asset farming
---

# Return Rates in Crypto Farming

Decentralized Finance \(DeFi\) has attracted a lot of attention in the recent years. A rather new and hot topic inside the world of DeFi is yield farming, which allows to earn passive income on token holdings using various investment strategies coded into smart contracts.  
There is an ever growing number of products labeled as yield farming. Many of those are advertised with an APY \(Annual Percentage Yield\) rate.  
We emphasize that in this context, the term APY is to be read with caution. In classical finance, it refers to the annual rate of return in an investment and is a predefined rate. Hence, an investor can be sure to obtain the corresponding interest in their investment. However, most farming products heavily rely on the investment strategy in volatile crypto markets. In contrast to APYs in lending and borrowing, rates published in farming are either APYs of past periods or projections based on past data.

## Pool Rates

Due to the complexity of this topic, our first step consists of publishing the pool rates as emitted in the smart contracts of various farming products. The basic idea of pool rates is pretty straightforward and can be understood as follows. Consider a farming pool of token $$X$$ and let $$A_X(i)$$ be the amount of token $$X$$ at block number $$i$$. Assume, farming starts at block number $$i_0$$, then we have the following equation

$$
A_X(i)=A_X(i_0)\cdot p_X(i)
$$

where $$p_X(i)$$ is the price corresponding to the pool of token $$X$$ . Initially, $$i=i_0$$ and hence $$p_X(i_0)=1$$. Motivated by this fact, we write

$$
p_X(i)=1+r_X(i).
$$

Here, $$r_X(i)$$ is the _pool rate_ corresponding to the pool of token $$X$$ .

As the initial supply $$A_X(i_0)$$ is constant, the price depends on the evolution of $$A_X(i)$$ and hence on the pool's investment strategy - if the number of tokens $$A_X(i)$$ increases with block number $$i$$, so will the pool rate. In other words, the return of an investment is determined by the change in the pool rate.  
The following example illustrates the relation between the pool rate and the return on an investment. We remark that in general, as for interest rates in traditional finance, differences of pool rates can be used for simple interest calculation as well as for compounded interest.

**Example:**  
Assume an investor puts an amount of $$100$$ tokens $$X$$ into an empty pool at block number $$10000000$$. With the notation from above this means $$i_0=10000000$$ and $$A_X(10000000)=100$$. Further assume that $$1000$$ blocks later, the investment strategy has increased the number of tokens in the pool to $$101$$ tokens, so $$A_X(10001000)=101$$. The return on the initial investment after these $$1000$$ blocks can be obtained by the difference of the pool rates at blocks $$10001000$$ and $$10000000$$ respectively. Indeed, we have $$p_X(10000000)=1$$ and thus $$r_X(10000000)=0$$. By the above equation for $$A_X(i)$$ we have $$p_X(10001000)=\frac{A_X(10001000)}{A_X(10000000)}=\frac{101}{100}$$ and thus $$r_X(10001000)=\frac{1}{100}$$. This yields

$$
r_X(10001000)-r_X(10000000)=\frac{1}{100}
$$

and hence a return rate of 1%.

Again, whether the amount of tokens in a pool increases or decreases solely depends on the investment strategy. The return rates for past time ranges can be computed by considering pool rate differences, such as done in the example above. In order to estimate future return rates, one can apply mathematical methods which allow for the estimation of future values based on past data. The simplest way of doing so is to fit a linear function to the data, which is also known under the name of \(linear\) regression. However, for most cases, such simple models do not yield good results for bigger time ranges. Mathematicians have tackled such problems since a long time and there is a wide range of techniques available. Nowadays, many of these are used under the name of _machine learning_.

## Total Debt Systems

A slightly different approach for farming is presented by other platforms such as Synthetix. Apart from constant fee rates on transactions such as trades,  the dynamical return rate depends on the  average return of all investments on the platform. The corresponding gains \(or losses\) for the investors are paid out periodically, for instance weekly. We will first explain the general mechanism and then illustrate it with an example.

Consider a platform with $$n$$ different pools associated to an investment strategy. In most cases, the investment strategy is short or long on a cryptocurrency, represented by a minted synthetic asset. Assume an investor makes an investment by buying from a pool of such a token $$X$$  worth $$p_0$$  US-Dollar \(it could be any other currency obviously\). She thereby increases the so-called system's _total debt_ from the previous value to $$D_0$$ \(also in US-Dollar\). Here, the total debt  is just the value of all investments across the platform. Now assume that after one farming period has elapsed, the value of her investment is $$p_1$$.  Furthermore, the system's total debt, i.e. the US-Dollar value of all investments together changed to $$D_1$$. The idea is that the system does not owe the investor her initial investment, but rather her initial proportion of the system's total debt, which is $$\frac{p_0}{D_0}$$. Hence, her actual personal debt $$D_{Inv}$$ is

$$
D_{Inv}=\frac{p_0}{D_0}\cdot D_1.
$$

In other words, the total debt is distributed among all investors according to their initial share of the total debt. In particular, everyone wins if $$D_1>D_0$$, that is, the total debt increases, and everyone looses if $$D_1<D_0$$. Furthermore, the investor's total gain $$G$$ after one farming period is given by her personal debt minus her initial investment

$$
G=D_{Inv}-p_0=p_0\left( \frac{D_1}{D_0}-1 \right).
$$

Let us quickly illustrate this equation before making the connection to the pool rate $$r$$. Assume the total debt has increased from $$D_0=1000$$  by $$10\%$$ to $$D_1=1100$$. Plugging into the above equation for the gain $$G$$ leads to $$G=p_0\frac{1}{10}$$ and hence, the investor makes a gain of $$10\%$$.  
In order to see the relation to the _pool rate_ $$r$$  from the previous section, we rewrite the personal debt $$D_{Inv}$$  in a somewhat artificial manner as

$$
D_{Inv}=p_0\left(1+\left(\frac{D_1}{D_0}-1\right)\right).
$$

And hence, the pool rate $$r$$ is given by

$$
r=\frac{D_1}{D_0}-1.
$$

It should come as no surprise that this leads to the pool rate $$r=0.1$$ for the above example.  
Let us now consider a slightly more comprehensive example.

**Example:**

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Step</b>
      </th>
      <th style="text-align:left">Investment 1 in $</th>
      <th style="text-align:left">Investment 2 in $</th>
      <th style="text-align:left">Investment 3 in $</th>
      <th style="text-align:left">Total Debt in $</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">
        <p>sBTC</p>
        <p>sETH</p>
      </td>
      <td style="text-align:left">
        <p>0</p>
        <p>100</p>
      </td>
      <td style="text-align:left">0
        <br />48</td>
      <td style="text-align:left">
        <p>2</p>
        <p>0</p>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">
        <p>sBTC</p>
        <p>sETH</p>
      </td>
      <td style="text-align:left">
        <p>0</p>
        <p>100</p>
      </td>
      <td style="text-align:left">
        <p>0</p>
        <p>48</p>
      </td>
      <td style="text-align:left">
        <p>4</p>
        <p>0</p>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
  </tbody>
</table>

Consider a protocol with just two pools, namely sBTC and sETH, and three investors as shown in the first row. The system's total debt is the sum of the investments. Assume that after one farming period the ETH price miraculously remains the same, and the price of BTC doubles, leading to the balances shown in row 2. In the third row, the personal debt for each investor is computed using the above formula for $$D_{Inv}$$ . In the last row, the corresponding total gains are computed. Although investor 3's strategy made a gain of $$100\%$$, the investor himself only made a gain of $$1,5\%$$, as he was the only one winning and his total gain is distributed among all investors. Furthermore, it is distributed according to the initial shares. As Investors 1 and 2 have far bigger initial shares, they get a large amount of Investment 3's gain, although their strategy did not win.  
Finally, we can easily compute the farming period's pool rate for this example, using the above formula

$$
r=\frac{152}{150}-1=0,01\bar{3}=1,\bar{3}\%.
$$

## Estimation of Returns from Total Rewards

Some platforms such as Loopring emit total system rewards. We remark that, given the staking pool's total balance $$D$$ and the total reward $$R$$, an average pool rate $$\tilde{r}$$ can be estimated as follows

$$
\tilde{r}=\frac{R}{D}.
$$

 

