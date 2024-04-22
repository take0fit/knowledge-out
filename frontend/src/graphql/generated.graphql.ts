import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Date: { input: any; output: any; }
};

export type Error = {
  __typename?: 'Error';
  code: Scalars['Int']['output'];
  message: Scalars['String']['output'];
};

export type Input = {
  __typename?: 'Input';
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  inputCategoryId: Scalars['Int']['output'];
  inputDetail?: Maybe<Scalars['String']['output']>;
  inputName: Scalars['String']['output'];
  outputs: Array<Output>;
  resourceId: Scalars['ID']['output'];
  updatedAt: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createInput: Input;
  createOutput: Output;
  createResource: Resource;
  createUser: User;
};


export type MutationCreateInputArgs = {
  inputCategoryId: Scalars['Int']['input'];
  inputDetail?: InputMaybe<Scalars['String']['input']>;
  inputName: Scalars['String']['input'];
  resourceId: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};


export type MutationCreateOutputArgs = {
  inputIds: Array<Scalars['String']['input']>;
  outputCategoryId: Scalars['Int']['input'];
  outputDetail?: InputMaybe<Scalars['String']['input']>;
  outputName: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};


export type MutationCreateResourceArgs = {
  resourceCategoryId: Scalars['Int']['input'];
  resourceDetail?: InputMaybe<Scalars['String']['input']>;
  resourceName: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};


export type MutationCreateUserArgs = {
  birthday?: InputMaybe<Scalars['Date']['input']>;
  nickname: Scalars['String']['input'];
};

export type Output = {
  __typename?: 'Output';
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  inputId: Scalars['ID']['output'];
  outputCategoryId: Scalars['Int']['output'];
  outputDetail?: Maybe<Scalars['String']['output']>;
  outputName: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['String']['output']>;
  hasNextPage: Scalars['Boolean']['output'];
  hasPreviousPage: Scalars['Boolean']['output'];
  startCursor?: Maybe<Scalars['String']['output']>;
};

export type Query = {
  __typename?: 'Query';
  input?: Maybe<Input>;
  inputsByUserId: Array<Input>;
  output?: Maybe<Output>;
  outputsByUserId: Array<Output>;
  resource?: Maybe<Resource>;
  resourcesByUserId: Array<Resource>;
  user?: Maybe<User>;
  users: Array<User>;
};


export type QueryInputArgs = {
  id: Scalars['ID']['input'];
};


export type QueryInputsByUserIdArgs = {
  userId: Scalars['ID']['input'];
};


export type QueryOutputArgs = {
  id: Scalars['ID']['input'];
};


export type QueryOutputsByUserIdArgs = {
  userId: Scalars['ID']['input'];
};


export type QueryResourceArgs = {
  id: Scalars['ID']['input'];
};


export type QueryResourcesByUserIdArgs = {
  userId: Scalars['ID']['input'];
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type Resource = {
  __typename?: 'Resource';
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  inputs: Array<Input>;
  resourceCategoryId: Scalars['Int']['output'];
  resourceDetail?: Maybe<Scalars['String']['output']>;
  resourceName: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type User = {
  __typename?: 'User';
  age?: Maybe<Scalars['Int']['output']>;
  birthday?: Maybe<Scalars['Date']['output']>;
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  nickname: Scalars['String']['output'];
  resources: Array<Resource>;
  updatedAt: Scalars['String']['output'];
};

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = { __typename?: 'Query', users: Array<{ __typename?: 'User', id: string, nickname: string, age?: number | null, birthday?: any | null }> };


export const GetUsersDocument = gql`
    query GetUsers {
  users {
    id
    nickname
    age
    birthday
  }
}
    `;

export function useGetUsersQuery(options?: Omit<Urql.UseQueryArgs<GetUsersQueryVariables>, 'query'>) {
  return Urql.useQuery<GetUsersQuery, GetUsersQueryVariables>({ query: GetUsersDocument, ...options });
};