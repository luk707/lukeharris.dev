module.exports = {
  siteMetadata: {
    title: `Luke Harris`,
  },
  plugins: [
    `gatsby-plugin-styled-components`,
    `gatsby-plugin-react-helmet`,
    `gatsby-transformer-remark`,
    {
      resolve: `gatsby-source-filesystem`,
      options: {
        path: `./src/content/`,
      },
    },
  ],
};
