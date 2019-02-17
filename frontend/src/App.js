import React from 'react';
import gql from 'graphql-tag';
import { Query } from 'react-apollo';

import './App.css';

const GET_HELLO_WORLD = gql`
  {
    hello
  }
`;

const App = () => (
  <Query query={GET_HELLO_WORLD}>
    {({ loading, error, data }) => {
      if (loading) {
        return <div>Loading ...</div>;
      }

      if (error) {
        return <p>Error :( {error.message}</p>;
      }

      return (
        <div>
          <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
      );
    }}
  </Query>
);

export default App;
