---
home: true
//heroImage: images/pixelated-logo.png
//actionText: Get Started →
//actionLink: /guide/getting-started.html
---

<DIAComponentHeader />

# Introduction into documentation
This page shows the most important features of the new documentation tool - so consider this a list of tools and assets you can pick, use, and modify is needed

# What is DIA? 
__DIA (Decentralised Information Asset) is a cross-chain, end-to-end, open-source data and oracle platform for Web3.__
The DIA platform enables the sourcing, validation and sharing of transparent and verified data feeds for traditional and digital financial applications. DIA’s institutional-grade data feeds cover asset prices, metaverse data, lending rates and more.
DIA’s data is directly sourced from a broad array of on-chain and off-chain sources at individual a trade-level. This allows DIA feeds to be fully customised with regards to the mix of sources and methodologies, resulting in tailor-made, high resilience feeds, setting a new paradigm for oracles.

# Markdown Table (VuePress styling)

| DATA ACCESS | CONTRIBUTE | KNOWLEDGE BASE |
| ----------- | ---------- | -------------- |
| Intro text Data access   | Introtext Contribute | Introtext Knowledge Base |
| link [Data access](doc_dir_data-access/ "title text link hello world data access!") | link [Contribute](doc_dir_contribute/ "title text link hello world data access!") | Link Knowledge Base |


# Image test
![asfasfasdf](images/somelogo.png)



# HTML Table (individual styling)
<div class="features">
  <div class="feature">
    <h2>DATA ACCESS</h2>
    <p>Intro text Data access</p>
    <a href="doc_dir_data-access/">Data access</a>
  </div>
  <div class="feature">
    <h2>CONTRIBUTE</h2>
    <p>Introtext Contribute</p>
    <a href="doc_dir_contribute/">link Contribute</a>
  </div>
  <div class="feature">
    <h2>KNOWLEDGE BASE</h2>
    <p>Introtext Knowledge Base</p>
    <a href="asdfasdfa">asdfasdf</a>
  </div>
</div>

::: tip
This is a tip
:::

::: warning
This is a warning
:::

::: danger
This is a dangerous warning
:::

::: danger CUSTOM TITLE
This is a dangerous warning with a custom title
:::

::: details
This is a details block, which does not work in IE / Edge
:::


# SampleComponentMarkdownContent
<SampleComponentMarkdownContent>

### Example component with markdown content 
- Any markdown
- you put here
> will be slotted in 

</SampleComponentMarkdownContent>

# SampleComponentJavascriptExecution (also parameter passing)
<SampleComponentJavascriptExecution display-text="display text" headlineText="hello world" />


# SampleComponentConditional Content
<SampleComponentCondition MyConditionalParameter='Peter'>
  test content only shown if MyConditionalParameter == Peter
</SampleComponentCondition>

<SampleComponentCondition MyConditionalParameter='xxx'>
  should not be shown
</SampleComponentCondition>

# SampleComponentForLoop For something like icon collections, can be passed as array of string ids
<SampleComponentForLoop MyList="ListItem1,ListItems2" />

# Syntax highlighting

``` js
var foo = function (bar) {
  return bar++;
};

console.log(foo(5));
```

<DIAComponentFooter />
