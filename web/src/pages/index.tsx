import React, { FC } from "react";
import { PageProps } from "gatsby";
import { Helmet } from "react-helmet";

const HomePage: FC<PageProps> = ({ path }) => (
  <>
    <Helmet>
      <meta
        name="go-import"
        content="lukeharris.dev git https://github.com/luk707/lukeharris.dev"
      />
    </Helmet>
    <h1>lukeharris.dev</h1>
  </>
);

export default HomePage;
