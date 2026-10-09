import js from '@eslint/js';
import tsPlugin from '@typescript-eslint/eslint-plugin';
import tsParser from '@typescript-eslint/parser';
import prettier from 'eslint-config-prettier';
import svelte from 'eslint-plugin-svelte';
import globals from 'globals';

export default [
	{ ignores: ['dist/', '.svelte-kit/', 'package/'] },
	js.configs.recommended,
	...tsPlugin.configs['flat/recommended'],
	...svelte.configs.recommended,
	prettier,
	{ languageOptions: { globals: { ...globals.browser, ...globals.node } } },
	{ files: ['**/*.svelte'], languageOptions: { parserOptions: { parser: tsParser } } }
];
