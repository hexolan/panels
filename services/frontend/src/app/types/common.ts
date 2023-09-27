// Auth
export type AuthData = {
  token: AuthToken;
  user: User;
}

export type AuthToken = {
  token_type: string;
  access_token: string;
  expires_in: number;
}

// Panel
export type Panel = {
  id: string;
  name: string;
  description: string;
  createdAt: string;
  updatedAt?: string;
}

// Post
export type Post = {
  id: string;
  panelId: string;
  authorId: string;
  title: string;
  content: string;
  createdAt: string;
  updatedAt?: string;
}

// Comment
export type Comment = {
  id: string;
  postId: string;
  authorId: string;
  message: string;
  createdAt: string;
  updatedAt?: string;
}

// User
export type User = {
  id: string;
  username: string;
  isAdmin: boolean;
  createdAt?: string;
  updatedAt?: string;
}