# Compounded Rates

In order to reduce the impact of outliers in daily fluctuations of financial markets, many financial products use simple or compounded averages of a given reference rate. For instance, on every business day the New York Federal Reserve Bank publishes the average rate of the Secured Overnight Financing Rate, compounded over the last 30 calendar days under the name of SOFR30.

## Standard Methodology

Here, we present the methodology of compounded rates that is used by large banks such as the New York Federal Reserve Bank \(compounded SOFR\) and the Bank of England \(compounded SONIA\).

Consider a unit investment over a time period of $$d_c$$ calendar days containing $$d_b\leq d_c$$ business days. Let $$r_i$$ be an interest rate that is published once every business day and assume that the business day convention is such that the year has $$N$$ days. Then, the compounded gain $$G$$ over the whole interest period results to

 $$\begin{equation} G= \prod_{i=1}^{d_b}\left( 1 + \frac{r_i \times n_i}{N} \right) \end{equation}$$ .

We remark that for constant interest rates $$r_i = r$$ and $$n_i = 1$$ this simplifies to the well-known compounded gain formula $$G=\left(1+\frac{r}{N}\right)^{d_b}$$.  
However, in general, the rate factor $$n_i$$ is an integer value that accounts for $$i$$ being a business day or a non-business day. More precisely, if $$i$$ is a business day followed by $$k$$ non-business days, we set $$n_i = k+1$$. For instance, if $$i$$ is followed by a business day, i.e., $$k=0$$, we have $$n_i=1$$. For a friday, which is usually followed by two non-business days, we would have $$n_i = 3$$.   
Now, in order to get the average interest $$I$$ from the compounded gain $$G$$, we subtract the original investment and normalize, thus obtaining

 $$I = \frac{N}{d_c}\left[\prod_{i=1}^{d_b}\left( 1 + \frac{r_i \times n_i}{N} \right) -1\right]$$ 

which is the formulation of compounded rates used by the FED and the BOE amongst others.

Sources:  
[https://www.newyorkfed.org/medialibrary/Microsites/arrc/files/2019/Users\_Guide\_to\_SOFR.pdf](https://www.newyorkfed.org/medialibrary/Microsites/arrc/files/2019/Users_Guide_to_SOFR.pdf)  
  
[https://www.bankofengland.co.uk/paper/2020/supporting-risk-free-rate-transition-through-the-provision-of-compounded-sonia](https://www.bankofengland.co.uk/paper/2020/supporting-risk-free-rate-transition-through-the-provision-of-compounded-sonia)

Link to API documentation:  
[https://docs.diadata.org/documentation/api-1/api-endpoints\#compounded-average](https://docs.diadata.org/documentation/api-1/api-endpoints#compounded-average)

## DIA Methodology

The methodology from the previous section has a special feature in that it mixes compounded and non-compounded rates. More precisely, investments are not compounded for weekends and holidays. This behaviour is reflected in the rate factor $$n_i$$. In the Index $$I_{DIA}$$ presented below,  investments are compounded over all calendar days in the respective interest period.

Consider a unit investment over a time period of $$d_c$$ calendar days. Let $$r_i$$ be an interest rate that is published once every business day and assume that the business day convention is such that the year has $$N$$ days. We define an interest rate $$\tilde{r}_i$$ such that $$\tilde{r}_i$$ coincides with $$r_i$$ on business days and is set to the rate of the previous business day if $$i$$ is a holiday or a weekend. In this straightforward manner we obtain an interest rate for all _calendar days_ and can now set

$$I_{DIA}=\frac{N}{dc}\left[\prod_{j=1}^{d_c}\left( 1 + \frac{\tilde{r}_j}{N} \right) -1\right].$$

Link to API documentation:  
Coming soon!



