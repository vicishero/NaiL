declare namespace NetParams {

  type AuthUserInfo = string;

  interface AuthUpdateUserPassword {
    /** 新密码 */
    password: string;
    /** 旧密码 */
    old_password: string;
  }

  interface SiteInfoReq {}

  interface IndexTrendsReq {
    page: number;
    page_size: number;
  }

  interface PostGetPost {
    id: string;
  }

  interface PostGetPosts {
    query: string | null;
    type: string;
    style: 'newest' | 'hots' | 'following' | 'search';
    page: number;
    page_size: number;
  }

  interface PostLockPost {
    id: string;
  }

  interface PostStickPost {
    id: string;
  }

  interface PostHighlightPost {
    id: string;
  }

  interface PostVisibilityPost {
    id: string;
    /** 可见性：0为公开，1为私密，2为好友可见 */
    visibility: import('@/utils/IEnum').VisibilityEnum;
  }

  interface PostGetPostStar {
    id: string;
  }

  interface PostPostStar {
    id: string;
  }

  interface PostGetPostCollection {
    id: string;
  }

  interface PostPostCollection {
    id: string;
  }

  interface PostGetTags {
    type: 'hot' | 'new' | 'follow' | 'pin' | 'hot_extral';
    num: number;
    extral_num?: number;
  }

  interface PostGetPostComments {
    id: string;
    style: 'default' | 'hots' | 'newest';
    page?: number;
    page_size?: number;
  }

  interface PostCreatePost {
    /** 帖子内容列表 */
    contents: Partial<Item.PostItemProps>[];
    /** 标签列表 */
    tags: string[];
    /** 艾特用户列表 */
    users: string[];
    /** 附件价格 */
    attachment_price: number;
    /** 可见性：0为公开，1为私密，2为好友可见 */
    visibility: import('@/utils/IEnum').VisibilityEnum;
  }

  interface PostDeletePost {
    id: string;
  }

  interface PostTweetCommentThumbs {
    tweet_id: string;
    comment_id: string;
  }

  interface PostTweetReplyThumbs {
    tweet_id: string;
    comment_id: string;
    reply_id: string;
  }

  interface PostCreateComment {
    /** 内容ID */
    post_id: string;
    /** 帖子内容列表 */
    contents: Partial<Item.CommentItemProps>[];
    /** 艾特用户列表 */
    users: string[];
  }

  interface PostDeleteComment {
    id: string;
  }

  interface PostHighlightComment {
    id: string;
  }

  interface PostCreateCommentReply {
    /** 艾特的用户UID */
    at_user_id: string;
    /** 回复的评论ID */
    comment_id: string;
    /** 回复内容 */
    content: string;
  }

  interface PostDeleteCommentReply {
    id: string;
  }

  interface PostStickTopic {
    topic_id: string;
  }

  interface PostPinTopic {
    topic_id: string;
  }

  interface PostFollowTopic {
    topic_id: string;
  }

  interface PostUnfollowTopic {
    topic_id: string;
  }
}
