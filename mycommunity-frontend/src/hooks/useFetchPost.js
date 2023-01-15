import api from "./api";
import { useMutation, useInfiniteQuery } from "react-query";

const FEED_ENDPOINT = "/api/v1/feed";

export const useFeed = (options = {}) => {
  return useInfiniteQuery(
    "feed",
    async ({ pageParam = 1, queryKey }) => {
      const { data } = await api.get(`${FEED_ENDPOINT}/${pageParam}`);
      return data.body;
    },
    {
      getNextPageParam: (page) => {
        return page.pagination.page === page.pagination.lastpage
          ? undefined
          : page.pagination.page + 1;
      },
      //   keepPreviousData: false,
    }
  );
};
