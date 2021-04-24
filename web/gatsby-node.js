const path = require("path");

exports.createPages = async ({ graphql, actions }) => {
  const { createPage } = actions;
  const result = await graphql(`
    query {
      allGoPackagesYaml {
        edges {
          node {
            id
            import
          }
        }
      }
    }
  `);

  result.data.allGoPackagesYaml.edges.forEach(({ node }) => {
    createPage({
      path: node.import,
      component: path.resolve(`./src/templates/go-package.tsx`),
      context: node,
    });
  });
};
