const { gitDescribe, gitDescribeSync } = require('git-describe');
process.env.VUE_APP_GIT_VERSION = gitDescribeSync()

module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  publicPath: '/web/',
  devServer: {
    proxy: {
      '/rpc': {
        target: 'http://localhost:22222/'
      },
      '/ws': {
        target: 'http://localhost:22222/',
        changeOrigin: true,
        ws: true,
      }
    }
  }
}
