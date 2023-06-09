import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import { terser } from 'rollup-plugin-terser';
import css from 'rollup-plugin-css-only';
import scss from 'rollup-plugin-scss';
import sass from 'sass';
import { existsSync, mkdirSync, writeFileSync } from 'fs';
import purify from 'purify-css';
import copy from 'rollup-plugin-copy';
import replace from '@rollup/plugin-replace';

const production = !process.env.ROLLUP_WATCH;
const buildDir = 'public/build';

function themeGenerator(theme) {
	return {
		input: `src/theme-${theme}.js`,
		output: {
			sourcemap: !production,
			format: 'iife',
			name: `appTheme_${theme}`,
			file: `${buildDir}/theme-${theme}.js`
		},
		plugins: [
			scss({
				includePaths: [
					'node_modules',
					'src'
				],
				runtime: sass,
				output: function (styles, styleNodes) {
					if (!existsSync(buildDir)) {
						mkdirSync(buildDir, { recursive: true });
					}
					writeFileSync(`${buildDir}/theme-${theme}.css`, styles);
					purify(
						[
							'node_modules/@hotwired/**/*.js',
							'node_modules/@sargassum-world/**/*.js',
							'src/**/*.js',
							'../templates/**/*.tmpl',
						],
						styles,
						{
							output: `${buildDir}/theme-${theme}.min.css`,
							minify: true,
							info: true,
							cleanCssOptions: {
								level: 1,
							},
						},
					);
				},
			}),
		],
		watch: {
			clearScreen: false
		}
	};
}

function bundleGenerator(type, appName, context) {
	return {
		input: `src/main-${type}.js`,
		output: {
			sourcemap: !production,
			format: 'iife',
			name: appName,
			file: `${buildDir}/bundle-${type}.js`
		},
		context,
		plugins: [
			// manually copy fontsource fonts, since rollup refuses to do it on its own
			copy({
				targets: [
					{ src: 'node_modules/@fontsource/*/files/*', dest: `${buildDir}/fonts` }
				]
			}),
			// we'll extract any component CSS out into
			// a separate file - better for performance
			css({ output: `${buildDir}/bundle-${type}.css` }),

			// If you have external dependencies installed from
			// npm, you'll most likely need these plugins. In
			// some cases you'll need additional configuration -
			// consult the documentation for details:
			// https://github.com/rollup/plugins/tree/master/packages/commonjs
			resolve({
				browser: true,
			}),
			commonjs(),

			// If we're building for production (npm run build
			// instead of npm run dev), minify
			production && terser()
		],
		watch: {
			clearScreen: false
		}
	};
}

export default [
	themeGenerator('eager', undefined),
	themeGenerator('light'),
	themeGenerator('dark'),
	bundleGenerator('eager', 'appEager'),
	bundleGenerator('deferred', 'app', 'window'),
];
