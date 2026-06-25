/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: [
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './public/index.html'
  ],
  theme: {
    extend: {
      colors: {
        'primary-fixed': '#ffdbcc',
        'on-secondary-container': '#004b65',
        'on-tertiary-container': '#003357',
        'tertiary-fixed': '#d0e4ff',
        'background': '#131313',
        'inverse-on-surface': '#313030',
        'surface-container-low': '#1c1b1b',
        'on-surface': '#e5e2e1',
        'surface': '#131313',
        'surface-container-high': '#2a2a2a',
        'tertiary-fixed-dim': '#9ccaff',
        'on-secondary-fixed': '#001e2b',
        'on-primary-fixed': '#351000',
        'inverse-primary': '#a04100',
        'on-secondary': '#003548',
        'error-container': '#93000a',
        'surface-dim': '#131313',
        'on-primary': '#561f00',
        'on-tertiary': '#003257',
        'secondary-container': '#00c1fd',
        'primary': '#ffb693',
        'surface-tint': '#ffb693',
        'outline-variant': '#5a4136',
        'on-primary-fixed-variant': '#7a3000',
        'surface-container-lowest': '#0e0e0e',
        'on-error-container': '#ffdad6',
        'surface-container': '#201f1f',
        'on-tertiary-fixed-variant': '#00497b',
        'surface-variant': '#353534',
        'tertiary-container': '#059eff',
        'on-error': '#690005',
        'primary-fixed-dim': '#ffb693',
        'secondary': '#8fd8ff',
        'on-primary-container': '#572000',
        'surface-bright': '#393939',
        'error': '#ffb4ab',
        'primary-container': '#ff6b00',
        'secondary-fixed': '#c2e8ff',
        'on-tertiary-fixed': '#001d35',
        'outline': '#a98a7d',
        'surface-container-highest': '#353534',
        'on-secondary-fixed-variant': '#004d67',
        'tertiary': '#9ccaff',
        'secondary-fixed-dim': '#75d1ff',
        'on-surface-variant': '#e2bfb0',
        'inverse-surface': '#e5e2e1',
        'on-background': '#e5e2e1'
      },
      borderRadius: {
        'DEFAULT': '0.25rem',
        'lg': '0.5rem',
        'xl': '0.75rem',
        'full': '9999px'
      },
      spacing: {
        'margin-desktop': '40px',
        'container-max': '1280px',
        'base': '8px',
        'margin-mobile': '16px',
        'gutter': '24px'
      },
      fontFamily: {
        'headline-lg': ['Montserrat', 'sans-serif'],
        'label-md': ['Inter', 'sans-serif'],
        'headline-md': ['Montserrat', 'sans-serif'],
        'headline-lg-mobile': ['Montserrat', 'sans-serif'],
        'label-sm': ['Inter', 'sans-serif'],
        'display': ['Montserrat', 'sans-serif'],
        'body-md': ['Inter', 'sans-serif'],
        'body-lg': ['Inter', 'sans-serif']
      },
      fontSize: {
        'headline-lg': ['32px', { lineHeight: '1.2', fontWeight: '700' }],
        'label-md': ['14px', { lineHeight: '1', letterSpacing: '0.05em', fontWeight: '600' }],
        'headline-md': ['20px', { lineHeight: '1.3', fontWeight: '600' }],
        'headline-lg-mobile': ['24px', { lineHeight: '1.2', fontWeight: '700' }],
        'label-sm': ['12px', { lineHeight: '1', fontWeight: '500' }],
        'display': ['48px', { lineHeight: '1.1', letterSpacing: '-0.02em', fontWeight: '800' }],
        'body-md': ['16px', { lineHeight: '1.5', fontWeight: '400' }],
        'body-lg': ['18px', { lineHeight: '1.6', fontWeight: '400' }]
      }
    }
  },
  plugins: []
}