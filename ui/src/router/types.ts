import { IndexRouteObject, NonIndexRouteObject } from 'react-router-dom';

type CustomRouteObject = {
  page: string;
  rules?: string[];
};

type IndexRouteNode = IndexRouteObject & CustomRouteObject;

interface NonIndexRouteNode extends NonIndexRouteObject, CustomRouteObject {
  children?: (IndexRouteNode | NonIndexRouteNode)[];
}

export type RouteNode = IndexRouteNode | NonIndexRouteNode;
