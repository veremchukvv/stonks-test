import React from 'react';

const Home = (props: {name: string}) => {
  if (!props.name) {
      return (
          <>
            Please sign in
            </>)
  } return (
            <>
            Welcome {props.name}
            </>)
};

export default Home;