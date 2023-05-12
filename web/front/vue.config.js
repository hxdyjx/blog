module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '华夏第一剑修'
      return args
    })
  },
  productionSourceMap: false
}
