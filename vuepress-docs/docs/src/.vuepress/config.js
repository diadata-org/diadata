const { description } = require('../../package')

module.exports = {
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#title
   */
  title: 'DIADATA new documentation',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#description
   */
  description: description,

  /**
   * Extra tags to be injected to the page HTML `<head>`
   *
   * ref：https://v1.vuepress.vuejs.org/config/#head
   */
  head: [
    ['meta', { name: 'theme-color', content: '#3eaf7c' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
  ],

  /**
   * Theme configuration, here is the default theme configuration for VuePress.
   *
   * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
   */
  themeConfig: {
    repo: '',
    editLinks: false,
    docsDir: '',
    editLinkText: '',
    lastUpdated: false,
    nav: [
      {
        text: 'Request Custom Data',
        link: 'https://medium.com/dia-insights/how-to-submit-a-dia-custom-delivery-request-cdr-in-5-minutes-6f88b0a4ca56',
      },
      {
        text: 'Developer Discord',
        link: 'https://discord.gg/zFmXtPFgQj'
      },
      {
        text: 'DAO Forum',
        link: 'https://dao.diadata.org/'
      },
      {
        text: 'Relative link test',
        link: '/doc_dir_contribute/'
      },
      {
        text: 'DiaData',
        link: 'https://www.diadata.org/'
      }
    ],
    /**
    sidebar: [
      ['/', 'Sidebar Get Started (see config.js)'],
      ['/doc_dir_data-access/', 
       {
          title: 'Data access',
          collapsable: true,
          children: [
            '',
            'doc_content_data-access',
          ]
        }
      ],
      ['/doc_dir_contribute/', 'Sidebar Contribute'],
      ['/guide/', 'old2']
    ],
    */
    sidebar: {
      '/doc_dir_data-access/': [
        {
          title: 'Data access',
          collapsable: true,
          children: [
            '',
            'doc_content_data-access',
          ]
        }
      ],
      '/doc_dir_contribute/': [
        {
          title: 'Contribute',
          collapsable: true,
          children: [
            '',
            'doc_content_contribute',
          ]
        }
      ]
    }
    
  },


  /**
   * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
   */
  plugins: [
    '@vuepress/plugin-back-to-top',
    '@vuepress/plugin-medium-zoom',
  ]
}
