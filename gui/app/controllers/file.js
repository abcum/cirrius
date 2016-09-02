import Ember from "ember";

export default Ember.Controller.extend({

    actions: {

        delete: function() {

            this.model.destroyRecord().then( () => {
                this.transitionToRoute('project.index');
            });

        },

        edit: function() {

			let path = window.prompt("File path:");

			if (Ember.isEmpty(path)) return;

    		if ( path.charAt(path.length-1) === '/' ) return;

    		if ( !path.match(/^[a-zA-Z. /]+$/g) ) return;

			if ( this.get('model.project.files').map(i => i.get('path')).includes(path) ) return;

            this.model.set('path', path);

        },

        /*setupController: function(controller, model) {

            this._super(controller, model);

            this.controllerFor('files').send('select', model);

         },*/

    }

});
