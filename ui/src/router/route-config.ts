import { RouteNode } from '@/router/types';

const routeConfig: RouteNode[] = [
  {
    path: '/',
    page: 'pages/Layout',
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
        rules: ['isLoginAndNormal'],
      },
      {
        path: 'posts/:qid/edit',
        page: 'pages/Questions/Ask',
        rules: ['isLoginAndNormal'],
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
      },
      {
        path: 'users/register',
        page: 'pages/Users/Register',
      },
      {
        path: 'users/account-recovery',
        page: 'pages/Users/AccountForgot',
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
      },
      {
        path: '/users/account-activation/failed',
        page: 'pages/Users/ActivationResult',
      },
      {
        path: '/users/confirm-new-email',
        page: 'pages/Users/ConfirmNewEmail',
      },
      {
        path: '/users/account-suspended',
        page: 'pages/Users/Suspended',
      },
      // for admin
      {
        path: 'admin',
        page: 'pages/Admin',
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
export default routeConfig;
