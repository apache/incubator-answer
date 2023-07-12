import { FC, useLayoutEffect } from 'react';
import { Helmet } from 'react-helmet-async';

import Color from 'color';

import { shiftColor, tintColor, shadeColor } from '@/utils';
import { themeSettingStore } from '@/stores';

const Index: FC = () => {
  const { theme, theme_config } = themeSettingStore((_) => _);
  let primaryColor;
  if (theme_config?.[theme]?.primary_color) {
    primaryColor = Color(theme_config[theme].primary_color);
  }
  const setThemeColor = () => {
    const themeMetaNode = document.querySelector('meta[name="theme-color"]');
    if (themeMetaNode) {
      const themeColor = primaryColor ? primaryColor.hex() : '#0033ff';
      themeMetaNode.setAttribute('content', themeColor);
    }
  };
  useLayoutEffect(() => {
    setThemeColor();
  }, [primaryColor]);

  return (
    <Helmet>
      {primaryColor && (
        <style>
          {`
              :root {
                --bs-blue: ${primaryColor.hex()};
                --bs-primary: ${primaryColor.hex()};
                --bs-primary-rgb: ${primaryColor.rgb().array().join(',')};
                --bs-link-color: ${primaryColor.hex()};
                --bs-link-hover-color: ${shiftColor(primaryColor, 0.8)};
              }
              .nav-pills {
                --bs-nav-pills-link-active-bg: ${primaryColor.hex()};
              }
              .btn-primary {
                --bs-btn-bg: ${primaryColor.hex()};
                --bs-btn-border-color: ${primaryColor.hex()};
                --bs-btn-hover-bg: ${tintColor(primaryColor, 0.85)};
                --bs-btn-hover-border-color: ${tintColor(primaryColor, 0.9)};
                --bs-btn-focus-shadow-rgb: ${shadeColor(primaryColor, 0.85)};
                --bs-btn-active-bg: ${tintColor(primaryColor, 0.8)};
                --bs-btn-active-border-color: ${tintColor(primaryColor, 0.9)};
                --bs-btn-disabled-bg: ${primaryColor.hex()};
                --bs-btn-disabled-border-color: ${primaryColor.hex()};
              }
              .btn-outline-primary {
                --bs-btn-color: ${primaryColor.hex()};
                --bs-btn-border-color: ${primaryColor.hex()};
                --bs-btn-hover-bg: ${primaryColor.hex()};
                --bs-btn-hover-border-color: ${primaryColor.hex()};
                --bs-btn-active-bg: ${primaryColor.hex()};
                --bs-btn-active-border-color: ${primaryColor.hex()};
                --bs-btn-disabled-color: ${primaryColor.hex()};
                --bs-btn-disabled-border-color: ${primaryColor.hex()};
              }
              .pagination {
                --bs-btn-color: ${primaryColor.hex()};
                --bs-pagination-active-bg: ${primaryColor.hex()};
                --bs-pagination-active-border-color: ${primaryColor.hex()};
              }
              .form-select:focus,
              .form-control:focus,
               .form-control.focus{
                box-shadow: 0 0 0 0.25rem ${primaryColor
                  .fade(0.75)
                  .string()} !important;
                border-color: ${tintColor(primaryColor, 0.5)} !important;
              }
              .form-check-input:checked {
                background-color: ${primaryColor.hex()};
                border-color: ${primaryColor.hex()};
              }
              .form-check-input:focus {
                border-color: ${tintColor(primaryColor, 0.5)};
                box-shadow: 0 0 0 0.25rem rgba(var(--bs-primary-rgb), .4);
              }
              .form-switch .form-check-input:focus {
                background-image: url("data:image/svg+xml,%3csvg xmlns=%27http://www.w3.org/2000/svg%27 viewBox=%27-4 -4 8 8%27%3e%3ccircle r=%273%27 fill=%27${tintColor(
                  primaryColor,
                  0.5,
                )}%27/%3e%3c/svg%3e");
              }
              .dropdown-menu {
                --bs-dropdown-link-active-bg: rgb(var(--bs-primary-rgb));
              }
              .link-primary {
                color: ${primaryColor.hex()}!important;
              }
              .link-primary:hover, .link-primary:focus {
                color: ${shadeColor(primaryColor, 0.8).hex()}!important;
              }
            `}
        </style>
      )}
    </Helmet>
  );
};

export default Index;
