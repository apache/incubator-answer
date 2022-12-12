import { LoaderFunctionArgs, RouteObject } from 'react-router-dom';

import { guard } from '@/utils';
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
  guard?: (args: LoaderFunctionArgs) => Promise<TGuardResult>;
}

const routes: RouteNode[] = [
  {
    path: '/',
    page: 'pages/Layout',
    guard: async () => {
      return guard.notForbidden();
    },
    children: [
      // question and answer
      {
        index: true,
        page: 'pages/Questions',
      },
      {
        path: 'questions',
        page: 'pages/Questions',
      },
      {
        path: 'questions/ask',
        page: 'pages/Questions/Ask',
        guard: async () => {
          return guard.activated();
        },
      },
      {
        path: 'questions/:qid',
        page: 'pages/Questions/Detail',
      },
      {
        path: 'questions/:qid/:slugPermalink',
        page: 'pages/Questions/Detail',
      },
      {
        path: 'questions/:qid/:slugPermalink/:aid',
        page: 'pages/Questions/Detail',
      },
      {
        path: 'posts/:qid/edit',
        page: 'pages/Questions/Ask',
        guard: async () => {
          return guard.activated();
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
          return guard.activated();
        },
      },
      // for users
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
          return guard.logged();
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
          const notLogged = guard.notLogged();
          if (notLogged.ok) {
            return notLogged;
          }
          return guard.notActivated();
        },
      },
      {
        path: 'users/register',
        page: 'pages/Users/Register',
        guard: async () => {
          const allowNew = guard.allowNewRegistration();
          if (!allowNew.ok) {
            return allowNew;
          }
          return guard.notLogged();
        },
      },
      {
        path: 'users/account-recovery',
        page: 'pages/Users/AccountForgot',
        guard: async () => {
          return guard.activated();
        },
      },
      {
        path: 'users/change-email',
        page: 'pages/Users/ChangeEmail',
      },
      {
        path: 'users/password-reset',
        page: 'pages/Users/PasswordReset',
      },
      {
        path: 'users/account-activation',
        page: 'pages/Users/ActiveEmail',
      },
      {
        path: 'users/account-activation/success',
        page: 'pages/Users/ActivationResult',
        guard: async () => {
          return guard.activated();
        },
      },
      {
        path: '/users/account-activation/failed',
        page: 'pages/Users/ActivationResult',
        guard: async () => {
          return guard.notActivated();
        },
      },
      {
        path: '/users/confirm-new-email',
        page: 'pages/Users/ConfirmNewEmail',
      },
      {
        path: '/users/account-suspended',
        page: 'pages/Users/Suspended',
        guard: async () => {
          return guard.forbidden();
        },
      },
      {
        path: '/posts/:qid/timeline',
        page: 'pages/Timeline',
        guard: async () => {
          return guard.logged();
        },
      },
      {
        path: '/posts/:qid/:aid/timeline',
        page: 'pages/Timeline',
        guard: async () => {
          return guard.logged();
        },
      },
      {
        path: '/tags/:tid/timeline',
        page: 'pages/Timeline',
        guard: async () => {
          return guard.logged();
        },
      },
      // for admin
      {
        path: 'admin',
        page: 'pages/Admin',
        guard: async () => {
          await guard.pullLoggedUser(true);
          return guard.admin();
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
            path: 'themes',
            page: 'pages/Admin/Themes',
          },
          {
            path: 'css-html',
            page: 'pages/Admin/CssAndHtml',
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
          {
            path: 'branding',
            page: 'pages/Admin/Branding',
          },
          {
            path: 'legal',
            page: 'pages/Admin/Legal',
          },
          {
            path: 'write',
            page: 'pages/Admin/Write',
          },
          {
            path: 'seo',
            page: 'pages/Admin/Seo',
          },
          {
            path: 'login',
            page: 'pages/Admin/Login',
          },
        ],
      },
      // for review
      {
        path: 'review',
        page: 'pages/Review',
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
  {
    path: '/',
    page: 'pages/Layout',
    children: [
      {
        page: 'pages/Legal',
        children: [
          {
            path: 'tos',
            page: 'pages/Legal/Tos',
          },
          {
            path: 'privacy',
            page: 'pages/Legal/Privacy',
          },
        ],
      },
    ],
  },
  {
    path: '/install',
    page: 'pages/Install',
  },
  {
    path: '/maintenance',
    page: 'pages/Maintenance',
  },
];
export default routes;
