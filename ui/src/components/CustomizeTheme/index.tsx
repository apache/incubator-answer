/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { FC, useLayoutEffect } from 'react';
import { Helmet } from 'react-helmet-async';

import Color from 'color';

import { shiftColor, tintColor, shadeColor } from '@/utils';
import { themeSettingStore } from '@/stores';
import { DEFAULT_THEME_COLOR } from '@/common/constants';

const Index: FC = () => {
  const { theme, theme_config } = themeSettingStore((_) => _);
  let primaryColor;
  if (theme_config?.[theme]?.primary_color) {
    primaryColor = Color(theme_config[theme].primary_color);
  }
  const setThemeColor = () => {
    const themeMetaNode = document.querySelector('meta[name="theme-color"]');
    if (themeMetaNode) {
      const themeColor = primaryColor
        ? primaryColor.hex()
        : DEFAULT_THEME_COLOR;
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
                --bs-link-color-rgb: ${primaryColor.rgb().array().join(',')};
                --bs-link-hover-color: ${shiftColor(primaryColor, 0.8).hex()};
                --bs-link-hover-color-rgb: ${shiftColor(primaryColor, 0.8)
                  .round()
                  .array()}
              }
              :root[data-bs-theme='dark'] {
                --bs-link-color: ${tintColor(primaryColor, 0.6).hex()};
                --bs-link-color-rgb: ${tintColor(primaryColor, 0.6)
                  .round()
                  .array()};
                --bs-link-hover-color: ${shiftColor(
                  tintColor(primaryColor, 0.6),
                  -0.8,
                ).hex()};
                --bs-link-hover-color-rgb: ${shiftColor(
                  tintColor(primaryColor, 0.6),
                  -0.8,
                )
                  .round()
                  .array()};
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
              .tag-selector-wrap--focus {
                box-shadow: 0 0 0 0.25rem ${primaryColor
                  .fade(0.75)
                  .string()} !important;
                border-color: ${tintColor(primaryColor, 0.5)} !important;
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
