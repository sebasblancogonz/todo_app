import React, { Fragment, Component } from 'react'

class TaskElement extends Component {
  constructor(props) {
    super(props)
  }


  checkMark(completed) {
    
    if (completed) {
      return (
        <Fragment>
          <div className="checkmark">
            <div className="stem"></div>
            <div className="kick"></div>
          </div>
        </Fragment>
      )
    }

    return null
  }
  render() {
    const task = this.props
    return (
      <div key={task._id} className="wrap-task">
        <div className={'dot ' + (task && task.completed ? 'success' : '')}>
          {this.checkMark(task.completed)}
        </div>

        <div className="task">
          <div className="taskTitle">
            <div className="title">
              <h2>{task.title}</h2>
            </div>
          </div>
          <div className="taskDesc">
            <p>{task.description}</p>
          </div>
        </div>
        <div className="icons">
          <i
            className="fa fa-check"
            aria-hidden="true"
            onClick={() => this.props.toggleTask(task)}
          ></i>
          <div className="divider"></div>
          <i
            className="fa fa-trash"
            aria-hidden="true"
            onClick={() => this.props.onDelete(task._id)}
          ></i>
        </div>
      </div>
    )
  }
}

export default TaskElement