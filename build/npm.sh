#!/usr/bin/env sh

folder="$(pwd)"

versions=10

packages=()

bow=()
bow+=('async:::dist/async.min.js')
bow+=('bluebird:::js/browser/bluebird.min.js')
bow+=('lodash:::dist/lodash.min.js')
bow+=('underscore:::underscore-min.js')

npm=()
npm+=('accounting:::accounting.min.js')
npm+=('chance:::dist/chance.min.js')
npm+=('classnames:::index.js')
npm+=('co:::index.js')
npm+=('cookie:::index.js')
npm+=('dateformat:::lib/dateformat.js')
npm+=('deepmerge:::index.js')
npm+=('events:::events.js')
npm+=('extend:::index.js')
npm+=('hashids:::dist/hashids.min.js')
npm+=('immutable:::dist/immutable.min.js')
npm+=('jsondiffpatch:::public/build/jsondiffpatch-full.min.js')
npm+=('jsonic:::jsonic-min.js')
npm+=('merge:::merge.min.js')
npm+=('moment:::min/moment.min.js')
npm+=('deep-extend:::lib/deep-extend.js')
npm+=('odiff:::dist/odiff.umd.js')
npm+=('once:::once.js')
npm+=('promiz:::promiz.min.js')
npm+=('q:::q.js')
npm+=('qs:::dist/qs.js')
npm+=('react:::dist/react-with-addons.min.js')
npm+=('rx:::dist/rx.all.min.js')
npm+=('semver:::semver.js')
npm+=('string:::dist/string.min.js')
npm+=('sugar:::dist/sugar.min.js')
npm+=('traverse:::index.js')
npm+=('validator:::validator.min.js')
npm+=('when:::dist/browser/when.min.js')
npm+=('wrappy:::wrappy.js')
npm+=('xtend:::immutable.js')

npm cache clean && bower cache clean

for index in "${bow[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then

            echo "--- $name"#"$version"

			cd $folder
			bower install $name"#"$version
			cd "$folder/bower_components/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js" || { echo 'File not found'; exit 1; }
			rm -rf "$folder/bower_components/$name"

		fi

	done

done

for index in "${npm[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then

            echo "--- $name"@"$version"

			cd $folder
			npm install $name"@"$version
			cd "$folder/node_modules/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js" || { echo 'File not found'; exit 1; }
			rm -rf "$folder/node_modules/$name"

		fi

	done

done

npm cache clean && bower cache clean
rm -rf "$folder/bower_components/"
rm -rf "$folder/node_modules/"

exit 0
