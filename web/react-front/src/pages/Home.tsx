import React from 'react';

const Home = (props: {name: string}) => {

    return (
        <div>
            {props.name ? 'Welcome ' + props.name : 'Please sign in' }
        </div>
    );
};

export default Home;