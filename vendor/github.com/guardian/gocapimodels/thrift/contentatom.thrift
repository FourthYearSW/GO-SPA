namespace java com.gu.contentatom.thrift
#@namespace scala com.gu.contentatom.thrift

include "atoms/quiz.thrift"
include "atoms/media.thrift"
include "atoms/explainer.thrift"
include "shared.thrift"

typedef string ContentAtomID

enum AtomType { 
  QUIZ = 0,
//VIEWPOINTS = 1, DEPRECATED 
  MEDIA = 2,
  EXPLAINER = 3
}

union AtomData {
  1: quiz.QuizAtom quiz
//2: viewpoints.ViewpointsAtom viewpoints DEPRECATED
  3: media.MediaAtom media
  4: explainer.ExplainerAtom explainer
}

struct ContentChangeDetails {

    /** the latest change to the content atom */
     1: optional shared.ChangeRecord lastModified

    /** the atom creation event */
    2: optional shared.ChangeRecord created

    /** the atom publication event (if published) */
    3: optional shared.ChangeRecord published
    /**
     * the revision number of the content.
     *
     * This value is incremented whenever content is written to the database and can be used to
     * ensure message ordering.
     */
    4: required i64 revision
}

struct Flags {
  1: optional bool suppressFurniture
}

struct Atom {
  1: required ContentAtomID id
  2: required AtomType atomType
  3: required list<string> labels // required, but may be empty
  4: required string defaultHtml
  5: required AtomData data       // the atom payload
  6: required ContentChangeDetails contentChangeDetails
  7: optional Flags flags
 }

enum EventType { UPDATE, TAKEDOWN }

struct ContentAtomEvent {
  1: required Atom atom
  2: required EventType eventType
  3: required shared.DateTime eventCreationTime
}
