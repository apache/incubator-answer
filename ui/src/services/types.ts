export type ReportType = 'question' | 'answer' | 'comment' | 'user';
export type ReportAction = 'close' | 'flag' | 'review';
export interface ReportParams {
  type: ReportType;
  action: ReportAction;
}

export interface TagBase {
  display_name: string;
  slug_name: string;
}

export interface SynonymsTag {
  display_name: string;
  slug_name: string;
  tag_id: string;
  tag?: string;
  original_text?: string;
  parsed_text?: string;
}
export interface TagInfo {
  tag_id: string;
  slug_name: string;
  display_name: string;
  original_text: string;
  parsed_text: string;
  follow_count: number;
  question_count: number;
  is_follower: boolean;
  member_actions;
  created_at?;
  updated_at?;
  main_tag_slug_name?: string;
  excerpt?;
}
export interface QuestionParams {
  title: string;
  content: string;
  html: string;
  tags: Tag[];
}

export interface Tag {
  display_name?: string;
  slug_name: string;
  main_tag_slug_name?: string;
  original_text?: string;
  parsed_text?: string;
}

export interface Paging {
  page: number;
  page_size: number;
}

export interface RecordResult extends Paging {
  list: any[];
  count: number;
}

export interface ListResult {
  count: number;
  list: any[];
}

export interface AnswerParams {
  content: string;
  html: string;
  question_id: string;
  id: string;
}

export interface LoginReqParams {
  e_mail: string;
  /** password */
  pass: string;
  captcha_id?: string;
  captcha_code?: string;
}

export interface RegisterReqParams extends LoginReqParams {
  name: string;
}

export interface ModifyPassReq {
  old_pass: string;
  pass: string;
}

export interface ModifyUserReq {
  display_name: string;
  username?: string;
  avatar: string;
  bio: string;
  bio_html?: string;
  location: string;
  website: string;
}

export interface UserInfoBase {
  username: string;
  rank: number;
  display_name: string;
  avatar: string;
  website: string;
  location: string;
  ip_info?: string;
}

export interface UserInfoRes {
  /** input name */
  avatar: string;
  id: number;
  username: string;
  rank: number;
  bio: string;
  bio_html: string;
  location: string;
  website: string;
  create_time?: string;
  display_name?: string;
  /** value = 1 active; value = 2 inactivated
   */
  mail_status: number;
  e_mail?: string;
  /** roles */
  is_admin?: true;
  [prop: string]: any;
}

export interface AvatarUploadReq {
  file: FormData;
}

export interface ImgCodeReq {
  captcha_id?: string;
  captcha_code?: string;
}

export interface ImgCodeRes {
  captcha_id: string;
  captcha_img: string;
  verify: boolean;
}

export interface PssRetReq extends ImgCodeReq {
  e_mail: string;
}

export interface CheckImgReq {
  action: 'login' | 'e_mail' | 'find_pass';
}

export interface NoticeSetReq {
  notice_switch: boolean;
}

export interface QuDetailRes {
  id: string;
  title: string;
  content: string;
  html: string;
  tags: any[];
  view_count: number;
  unique_view_count?: number;
  answer_count: number;
  favorites_count: number;
  follow_counts: 0;
  accepted_answer_id: string;
  last_answer_id: string;
  create_time: string;
  update_time: string;
  user_info: User;
  answered: boolean;
  collected: boolean;

  [prop: string]: any;
}

export interface AnswersReq extends Paging {
  order?: 'default' | 'updated';
  question_id: string;
}

export interface User {
  username: string;
  rank: string;
  display_name: string;
  avatar: string;
  website: string;
  location: string;

  [prop: string]: any;
}

export interface AnswerContent {
  id: string;
  question_id: string;
  content: string;
  html: string;
  create_time: string;
  update_time: string;
  user_info: User;

  [prop: string]: any;
}

export interface AnswerRes {
  count: number;
  list: AnswerContent[];
  page?: number;
  total_page?: number;
}

export interface PostAnswerReq {
  content: string;
  html: string;
  question_id: string;
}

export interface PageUser {
  id;
  displayName;
  userName?;
  avatar_url?;
}

export interface LangsType {
  label: string;
  value: string;
}

/**
 * @description interface for Question
 */
export type QuestionOrderBy =
  | 'newest'
  | 'active'
  | 'frequent'
  | 'score'
  | 'unanswered';

export interface QueryQuestionsReq extends Paging {
  order: QuestionOrderBy;
  tags?: string[];
}

export type AdminQuestionStatus = 'available' | 'closed' | 'deleted';

export type AdminContentsFilterBy = 'normal' | 'closed' | 'deleted';

export interface AdminContentsReq extends Paging {
  status: AdminContentsFilterBy;
}

/**
 * @description interface for Answer
 */
export type AdminAnswerStatus = 'available' | 'deleted';

/**
 * @description interface for Users
 */
export type UserFilterBy = 'all' | 'inactive' | 'suspended' | 'deleted';

/**
 * @description interface for Flags
 */
export type FlagStatus = 'pending' | 'completed';
export type FlagType = 'all' | 'question' | 'answer' | 'comment';
export interface AdminFlagsReq extends Paging {
  status: FlagStatus;
  object_type: FlagType;
}

/**
 * @description interface for Admin Settings
 */
export interface AdminSettingsGeneral {
  name: string;
  short_description: string;
  description: string;
}

export interface AdminSettingsInterface {
  logo: string;
  language: string;
  theme: string;
}

export interface SiteSettings {
  general: AdminSettingsGeneral;
  interface: AdminSettingsInterface;
}
/**
 * @description interface for Activity
 */
export interface FollowParams {
  is_cancel: boolean;
  object_id: string;
}

/**
 * @description search request params
 */
export interface SearchParams {
  q: string;
  order: string;
  page: number;
  size?: number;
}

/**
 * @description search response data
 */
export interface SearchResItem {
  object_type: string;
  object: {
    id: string;
    title: string;
    excerpt: string;
    created_at: number;
    user_info: UserInfoBase;
    vote_count: number;
    answer_count: number;
    accepted: boolean;
    tags: TagBase[];
  };
}
export interface SearchRes {
  count: number;
  extra: any;
  list: SearchResItem[];
}
