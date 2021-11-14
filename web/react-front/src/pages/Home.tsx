import React, {useContext} from 'react';
import {AuthContext} from "../context/authContext";
import PortfolioList from "../components/PortfolioList";

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
                <PortfolioList/>
            </>)
};

export default Home;