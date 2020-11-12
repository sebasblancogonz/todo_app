import React, { Fragment, Component } from 'react'
import { AddTask, TaskList } from '../../Task'
import Header from '../Header'

class Home extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <Fragment>
        <Header />
        <div className="tasks">
          <AddTask/>
          <TaskList/>
        </div>
      </Fragment>
    )
  }
}

export default Home