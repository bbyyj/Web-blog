
module.exports = {
    devServer: {              //设置本地域名
        host: "localhost",
        port: 8082,
        proxy: {          //设置代理解决跨域问题
            "/": {
                // target: "http://localhost:8080",  
                target: "http://3b0142cb.r10.cpolar.top", 
                  //要跨域的域名
                changeOrigin: true                  //是否开启跨域
            }
        }
    },
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/'
}
