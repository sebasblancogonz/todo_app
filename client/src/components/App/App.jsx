import React, { Component } from 'react'
import { BrowserRouter, withRouter, Switch, Route } from 'react-router-dom'

import { Home } from '..'
import '../../styles/styles.css'

const App = () => (
  <BrowserRouter>
    <Switch>
      <Route exact path="/" component={Home} />
    </Switch>
  </BrowserRouter>
)

export default App;
