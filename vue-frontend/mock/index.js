export default [
  {
    url: "/api/login",
    method: "post",
    response: ({ query }) => {
      return {
        token: "jjjjjjjjjjjjjjj",
      };
    },
  },
  {
    url: "/api/logout",
    method: "post",
    response: ({ query }) => {
      return {
        token: "jjjjjjjjjjjjjjj",
      };
    },
  },
  {
    url: "/api/user/info",
    method: "get",
    response: ({ query }) => {
      return {
        email: "jjjjjjjjjjjjjjj",
        name: "xxxxxx",
        avatar: "https://api.multiavatar.com/Binx Bond.svg",
      };
    },
  },
];
