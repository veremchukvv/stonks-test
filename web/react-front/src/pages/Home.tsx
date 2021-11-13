import React, {useContext} from 'react';
import {AuthContext} from "../context/authContext";

const Home = () => {
    const auth = useContext(AuthContext)
  if (!auth.isAuthenticated) {
      return (
          <>
            Please sign in
            </>)
  } return (
            <>
            Welcome {auth.userName}
            </>)
};

export default Home;