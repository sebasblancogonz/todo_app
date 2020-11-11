const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const Dotenv = require('dotenv-webpack');
const envPath = process.env.NODE_ENV ? `.env.${process.env.NODE_ENV}` : '.env';
module.exports = {
    
    entry: [
        'babel-polyfill',
        './src/index.js',
    ],

    output: {
        publicPath: '/',
        filename: 'main.js'
    },

    resolve: {
        extensions: ['.js', '.jsx'],
    },

    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader'],
            },

            {
                test: /\.(jpe?g|png|gif|svg)$/i,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: 'src/img/[name].[ext]',
                        outputPath: 'img/',
                    },
                },
            },

            {
                test: /\.(sa|sc|c)ss$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: [{ loader: 'css-loader'}, 'sass-loader'],
                }),
            },
            {
                test: /\.html$/,
                use: {
                    loader: 'html-loader',
                    options: {
                        minimize: true,
                    },
                },
            },
            {
                test: /\.(otf|ttf|eot|woff|woff2)$/,
                loader: 'file-loader',
                options: {
                    name: '[name].[ext]',
                    outputPath: 'fonts/',
                },
            },
        ]
    },

    plugins: [
        new ExtractTextPlugin({ filename: '[name].css' }),
        new HtmlWebpackPlugin({
            template: './resources/index.html',
            filename: 'index.html',
            hash: true,
        }),
        new Dotenv({
            path: envPath
        })
    ],

    devServer: {
        historyApiFallback: true,
        publicPath: '/',
        contentBase: './src',
    },
};