import React, { useContext, useState } from "react";
import api from "./api";
import { useMutation, useInfiniteQuery } from "react-query";
import { useQueryClient } from "react-query";

const NEWPOST_ENDPOINT = "/api/v1/newpost";

export const useNewPost = () => {
  const mutation = useMutation({
    mutationFn: async (data) => {
      return api
        .post(NEWPOST_ENDPOINT, data)
        .then((res) => {})
        .catch((err) => console.log(err));
    },
  });
  return mutation;
};
