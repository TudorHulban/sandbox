module main

import vweb

struct App {
	vweb.Context
}

fn main() {
	app := App{}
	vweb.run(app, 3000)
}

['/']
pub fn (mut app App) index() vweb.Result {
	return app.text('alive')
}