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
        target: 'http://localhost:22222/'
      }
    }
  }
}
