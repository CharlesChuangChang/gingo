const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer:{
    port:8080,
    host:'localhost',
    open:true,//配置浏览器自动访问
    proxy: {
      '/api': {  //   若请求的前缀不是这个'/api'，那请求就不会走代理服务器
        target: 'http://localhost:7000',  //这里写路径 
        pathRewrite: { '^/api': '' }, //将所有含/api路径的，去掉/api转发给服务器
        ws: true,  //用于支持websocket
        changeOrigin: true   //用于控制请求头中的host值
      },
    }
  }
})