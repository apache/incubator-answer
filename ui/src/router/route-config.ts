import { RouteObject } from 'react-router-dom';

export interface RouteNode extends RouteObject {
  page: string;
  children?: RouteNode[];
  rules?: string[];
}
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
        page: 'pages/QuestionDetail',
      },
      {
        path: 'questions/:qid/:aid',
        page: 'pages/QuestionDetail',
      },
      {
        path: 'questions/ask',
        page: 'pages/Ask',
        rules: ['isLoginAndNormal'],
      },
      {
        path: 'posts/:qid/edit',
        page: 'pages/Ask',
        rules: ['isLoginAndNormal'],
      },
      {
        path: 'posts/:qid/:aid/edit',
        page: 'pages/EditAnswer',
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
        page: 'pages/TagDetail',
      },
      {
        path: 'tags/:tagName/info',
        page: 'pages/TagInfo',
      },
      {
        path: 'tags/:tagId/edit',
        page: 'pages/EditTag',
      },
      // users
      {
        path: 'users/:username',
        page: 'pages/Personal',
      },
      {
        path: 'users/:username/:tabName',
        page: 'pages/Personal',
      },
      {
        path: 'users/settings',
        page: 'pages/Settings',
        children: [
          {
            index: true,
            page: 'pages/Settings/Profile',
          },
          {
            path: 'profile',
            page: 'pages/Settings/Profile',
          },
          {
            path: 'notify',
            page: 'pages/Settings/Notification',
          },
          {
            path: 'account',
            page: 'pages/Settings/Account',
          },
          {
            path: 'interface',
            page: 'pages/Settings/Interface',
          },
        ],
      },
      {
        path: 'users/notifications/:type',
        page: 'pages/Notifications',
      },
      {
        path: 'users/login',
        page: 'pages/Login',
      },
      {
        path: 'users/register',
        page: 'pages/Register',
      },
      {
        path: 'users/account-recovery',
        page: 'pages/AccountForgot',
      },
      {
        path: 'users/password-reset',
        page: 'pages/PasswordReset',
      },
      {
        path: 'users/account-activation',
        page: 'pages/ActiveEmail',
      },
      {
        path: 'users/account-activation/success',
        page: 'pages/ActivationResult',
      },
      {
        path: '/users/account-activation/failed',
        page: 'pages/ActivationResult',
      },
      {
        path: '/users/confirm-new-email',
        page: 'pages/ConfirmNewEmail',
      },
      {
        path: '/users/account-suspended',
        page: 'pages/Suspended',
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
