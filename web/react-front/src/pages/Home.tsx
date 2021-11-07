import React from 'react';

const Home = (props: {name: string}) => {
    console.log(props.name)

    return (
        <div>
            {props.name ? 'Welcome ' + props.name : 'Please sign in' }
        </div>
    );
};

export default Home;