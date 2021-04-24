const path = require("path");

exports.createPages = async ({ graphql, actions }) => {
  const { createPage } = actions;
  const result = await graphql(`
    query {
      allMarkdownRemark {
        edges {
          node {
            id
            html
            parent {
              ... on File {
                name
                relativePath
              }
            }
          }
        }
      }
    }
  `);

  result.data.allMarkdownRemark.edges.forEach(({ node }) => {
    const [type] = node.parent.relativePath.split("/");
    switch (type) {
      case "go-packages":
        createPage({
          path: node.parent.name,
          component: path.resolve(`./src/templates/go-package.tsx`),
          context: {
            id: node.id,
          },
        });
        return;
      default:
        return;
    }
  });
};
