import Button from '@mui/material/Button'
import { Meta } from '@storybook/react'
import { fn } from '@storybook/test'

const meta: Meta<typeof Button> = {
  title: 'UI/Button',
  component: Button,
  args: {
    onClick: fn(),
  },
}

export default meta

function Template(args: any) {
  return <Button {...args}>Text</Button>
}

export const Default = {
  render: Template,
  args: {},
}
