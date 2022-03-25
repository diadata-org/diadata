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
        text: 'Introduction (Guide, edited old config)',
        link: '/guide/',
      },
      {
        text: 'Doc Introduction (new documentation content)',
        link: '/doc_dir_introduction/'
      },
      {
        text: 'Data Access (Config)',
        link: '/config/'
      },
      {
        text: 'VuePress',
        link: 'https://v1.vuepress.vuejs.org'
      },
      {
        text: 'DiaData',
        link: 'https://www.diadata.org/'
      }
    ],
    sidebar: {
      '/',
      '/doc_dir_introduction/': [
        {
          title: 'sitebar config for root dir, see config.js',
          collapsable: false,
          children: [
            '',
            'doc_content_introduction',
          ]
        }
      ],
      '/guide/': [
        {
          title: 'Introduction (Guide)',
          collapsable: false,
          children: [
            '',
            'using-vue',
          ]
        }
      ],
      '/config/': [
        {
          title: 'sidebare test vueconfig',
          collapsable: false,
          children: [
            '',
            'test',
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
