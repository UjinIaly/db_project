package impl

import (
	"db_project/internal/models"
	"db_project/internal/repositories"
	"db_project/internal/usecases"
	"db_project/pkg/errors"
)

type PostUseCaseImpl struct {
	postRepository   repositories.PostRepository
	userRepository   repositories.UserRepository
	threadRepository repositories.ThreadRepository
	forumRepository  repositories.ForumRepository
}

func CreatePostUseCase(
	postRepository repositories.PostRepository,
	userRepository repositories.UserRepository,
	threadRepository repositories.ThreadRepository,
	forumRepository repositories.ForumRepository,
) usecases.PostUseCase {
	return &PostUseCaseImpl{
		postRepository:   postRepository,
		userRepository:   userRepository,
		threadRepository: threadRepository,
		forumRepository:  forumRepository,
	}
}

func (postUseCase *PostUseCaseImpl) Get(postID int64, relatedData *[]string) (postFull *models.PostFull, err error) {
	postFull = new(models.PostFull)
	var post *models.Post
	post, err = postUseCase.postRepository.GetByID(postID)
	if err != nil {
		err = errors.ErrPostNotFound
	}
	postFull.Post = post

	for _, data := range *relatedData {
		switch data {
		case "user":
			var author *models.User
			author, err = postUseCase.userRepository.GetByNickname(postFull.Post.Author)
			if err != nil {
				err = errors.ErrUserNotFound
			}
			postFull.Author = author
		case "forum":
			var forum *models.Forum
			forum, err = postUseCase.forumRepository.GetBySlug(postFull.Post.Forum)
			if err != nil {
				err = errors.ErrForumNotExist
			}
			postFull.Forum = forum
		case "thread":
			var thread *models.Thread
			thread, err = postUseCase.threadRepository.GetByID(postFull.Post.Thread)
			if err != nil {
				err = errors.ErrThreadNotFound
			}
			postFull.Thread = thread
		}
	}
	return
}

func (postUseCase *PostUseCaseImpl) Update(post *models.Post) (err error) {
	oldPost, err := postUseCase.postRepository.GetByID(post.ID)
	if err != nil {
		err = errors.ErrThreadNotFound
		return
	}

	if post.Message != "" {
		if oldPost.Message != post.Message {
			oldPost.IsEdited = true
		}
		oldPost.Message = post.Message

		err = postUseCase.postRepository.Update(oldPost)
		if err != nil {
			return
		}
	}

	*post = *oldPost

	return
}
