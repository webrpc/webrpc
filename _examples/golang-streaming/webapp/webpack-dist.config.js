const path = require('path')
const fs = require('fs')
const webpack = require('webpack')
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const main = [
  // 'webpack-dev-server/client?http://0.0.0.0:4444',
  // 'webpack/hot/only-dev-server',
  'whatwg-fetch',
  './src/index.ts'
]

module.exports = {
  context: process.cwd(), // to automatically find tsconfig.json
  entry: {
    main: main
  },
  optimization: {
    namedModules: true
  },
  plugins: [
    // new webpack.HotModuleReplacementPlugin(),
    new ForkTsCheckerWebpackPlugin({
      tslint: true,
      checkSyntacticErrors: true,
      watch: ['./src'] // optional but improves performance (fewer stat calls)
    }),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('development'),
    }),
    new HtmlWebpackPlugin({
      inject: true,
      template: 'src/index.html'
    })
  ],
  module: {
    rules: [{
      test: /.tsx?$/,
      use: [{
        loader: 'ts-loader',
        options: {
          transpileOnly: true
        }
      }],
      exclude: path.resolve(process.cwd(), 'node_modules'),
      include: [
        path.resolve(process.cwd(), 'src'),
      ]
    }, {
      test: /\.(jpe?g|png|gif|svg)$/i,
      use: [{
        loader: 'url-loader',
        options: {
          limit: 8192000
        }
      }]
    }]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.png', '.jpg']
  }//,
  // devtool: 'inline-source-map',
  // devServer: {
  //   host: '0.0.0.0',
  //   port: 4444,
  //   open: false,
  //   hot: true,
  //   historyApiFallback: true,
  //   stats: 'errors-only',
  //   disableHostCheck: true,
  //   contentBase: path.resolve(process.cwd(), 'src/public'),
  //   publicPath: '/'
  // }
}
