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

import { FC, ReactNode, useEffect, useState } from 'react';
import { useNavigate, useLoaderData } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc, TGuardResult } from '@/utils/guard';

import RouteErrorBoundary from './RouteErrorBoundary';

const RouteGuard: FC<{
  children: ReactNode;
  onEnter: TGuardFunc;
  path?: string;
  page?: string;
}> = ({ children, onEnter, path, page }) => {
  const navigate = useNavigate();
  const loaderData = useLoaderData();
  const [gk, setKeeper] = useState<TGuardResult>({
    ok: true,
  });
  const [gkError, setGkError] = useState<TGuardResult['error']>();
  const applyGuard = () => {
    if (typeof onEnter !== 'function') {
      return;
    }

    const gr = onEnter({
      loaderData,
      path,
      page,
    });

    setKeeper(gr);
    if (
      gr.ok === false &&
      gr.error?.code &&
      /403|404|50X/i.test(gr.error.code.toString())
    ) {
      setGkError(gr.error);
      return;
    }
    setGkError(undefined);
    if (gr.redirect) {
      floppyNavigation.navigate(gr.redirect, {
        handler: navigate,
        options: { replace: true },
      });
    }
  };
  useEffect(() => {
    /**
     * By detecting changes to location.href, many unnecessary tests can be avoided
     */
    applyGuard();
  }, [window.location.href]);

  let asOK = gk.ok;
  if (gk.ok === false && gk.redirect) {
    /**
     * It is possible that the route guard verification fails
     *    but the current page is already the target page for the route guard jump
     * This should render `children`!
     */
    asOK = floppyNavigation.equalToCurrentHref(gk.redirect);
  }
  return (
    <>
      {asOK ? children : null}
      {gkError ? <RouteErrorBoundary errCode={gkError.code as string} /> : null}
    </>
  );
};

export default RouteGuard;
