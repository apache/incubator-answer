import { RouteObject } from 'react-router-dom';

import { Guard } from '@/utils';
import type { TGuardResult } from '@/utils/guard';

export interface RouteNode extends RouteObject {
  page: string;
  children?: RouteNode[];
  /**
   * a method to auto guard route before route enter
   * if the `ok` field in guard returned `TGuardResult` is true,
   * it means the guard passed then enter the route.
   * if guard returned the `TGuardResult` has `redirect` field,
   * then auto redirect route to the `redirect` target.
   */
  guard?: () => Promise<TGuardResult>;
}

const routes: RouteNode[] = [
  {
    path: '/',
    page: 'pages/Layout',
    guard: async () => {
      return Guard.notForbidden();
    },
    children: [
      // question and answer
      {
        index: true,
        page: 'pages/Questions',
      },
      {
        path: 'questions',
        index: true,
        page: 'pages/Questions',
      },
      {
        path: 'questions/:qid',
        page: 'pages/Questions/Detail',
      },
      {
        path: 'questions/:qid/:aid',
        page: 'pages/Questions/Detail',
      },
      {
        path: 'questions/ask',
        page: 'pages/Questions/Ask',
        guard: async () => {
          return Guard.activated();
        },
      },
      {
        path: 'posts/:qid/edit',
        page: 'pages/Questions/Ask',
        guard: async () => {
          return Guard.activated();
        },
      },
      {
        path: 'posts/:qid/:aid/edit',
        page: 'pages/Questions/EditAnswer',
      },
      {
        path: '/search',
        page: 'pages/Search',
      },
      // tags
      {
        path: 'tags',
        page: 'pages/Tags',
      },
      {
        path: 'tags/:tagName',
        page: 'pages/Tags/Detail',
      },
      {
        path: 'tags/:tagName/info',
        page: 'pages/Tags/Info',
      },
      {
        path: 'tags/:tagId/edit',
        page: 'pages/Tags/Edit',
        guard: async () => {
          return Guard.activated();
        },
      },
      // users
      {
        path: 'users/:username',
        page: 'pages/Users/Personal',
      },
      {
        path: 'users/:username/:tabName',
        page: 'pages/Users/Personal',
      },
      {
        path: 'users/settings',
        page: 'pages/Users/Settings',
        guard: async () => {
          return Guard.logged();
        },
        children: [
          {
            index: true,
            page: 'pages/Users/Settings/Profile',
          },
          {
            path: 'profile',
            page: 'pages/Users/Settings/Profile',
          },
          {
            path: 'notify',
            page: 'pages/Users/Settings/Notification',
          },
          {
            path: 'account',
            page: 'pages/Users/Settings/Account',
          },
          {
            path: 'interface',
            page: 'pages/Users/Settings/Interface',
          },
        ],
      },
      {
        path: 'users/notifications/:type',
        page: 'pages/Users/Notifications',
      },
      {
        path: 'users/login',
        page: 'pages/Users/Login',
        guard: async () => {
          const notLogged = Guard.notLogged();
          if (notLogged.ok) {
            return notLogged;
          }
          return Guard.notActivated();
        },
      },
      {
        path: 'users/register',
        page: 'pages/Users/Register',
        guard: async () => {
          return Guard.notLogged();
        },
      },
      {
        path: 'users/account-recovery',
        page: 'pages/Users/AccountForgot',
        guard: async () => {
          return Guard.activated();
        },
      },
      {
        path: 'users/change-email',
        page: 'pages/Users/ChangeEmail',
        // TODO: guard this (change email when user not activated) ?
      },
      {
        path: 'users/password-reset',
        page: 'pages/Users/PasswordReset',
        guard: async () => {
          return Guard.activated();
        },
      },
      {
        path: 'users/account-activation',
        page: 'pages/Users/ActiveEmail',
        guard: async () => {
          const notActivated = Guard.notActivated();
          if (notActivated.ok) {
            return notActivated;
          }
          return Guard.notLogged();
        },
      },
      {
        path: 'users/account-activation/success',
        page: 'pages/Users/ActivationResult',
        guard: async () => {
          return Guard.activated();
        },
      },
      {
        path: '/users/account-activation/failed',
        page: 'pages/Users/ActivationResult',
        guard: async () => {
          return Guard.notActivated();
        },
      },
      {
        path: '/users/confirm-new-email',
        page: 'pages/Users/ConfirmNewEmail',
        //  TODO: guard this
      },
      {
        path: '/users/account-suspended',
        page: 'pages/Users/Suspended',
        guard: async () => {
          return Guard.forbidden();
        },
      },
      // for admin
      {
        path: 'admin',
        page: 'pages/Admin',
        guard: async () => {
          await Guard.pullLoggedUser(true);
          return Guard.admin();
        },
        children: [
          {
            index: true,
            page: 'pages/Admin/Dashboard',
          },
          {
            path: 'dashboard',
            page: 'pages/Admin/Dashboard',
          },
          {
            path: 'answers',
            page: 'pages/Admin/Answers',
          },
          {
            path: 'flags',
            page: 'pages/Admin/Flags',
          },
          {
            path: 'general',
            page: 'pages/Admin/General',
          },
          {
            path: 'interface',
            page: 'pages/Admin/Interface',
          },
          {
            path: 'questions',
            page: 'pages/Admin/Questions',
          },
          {
            path: 'users',
            page: 'pages/Admin/Users',
          },
          {
            path: 'users/:user_id',
            page: 'pages/Admin/UserOverview',
          },
          {
            path: 'smtp',
            page: 'pages/Admin/Smtp',
          },
        ],
      },
      {
        path: '*',
        page: 'pages/404',
      },
      {
        path: '50x',
        page: 'pages/50X',
      },
    ],
  },
];
export default routes;
